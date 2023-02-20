package notify

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ cron.Job = (*NotifyJob)(nil)
var uniqueJob = sync.Map{}

type NotifyJob struct {
	data        *model.Notify                    // 数据
	channel     *model.Channel                   // 渠道
	completeJob func(*model.Notify)              // 完成Job
	done        func(*model.Notify, error) error // 回调函数
}

// NewNotifyJob 创建Job
func NewNotifyJob(data *model.Notify, channel *model.Channel, completeJobFunc func(*model.Notify), doneFunc func(*model.Notify, error) error) *NotifyJob {
	return &NotifyJob{
		data:        data,
		channel:     channel,
		done:        doneFunc,
		completeJob: completeJobFunc,
	}
}

func (n *NotifyJob) Run() {
	// 判断是否执行中
	if _, ok := uniqueJob.Load(n.data.ID); ok {
		return
	}
	// 加入执行中
	uniqueJob.Store(n.data.ID, true)
	// 移除执行中
	defer uniqueJob.Delete(n.data.ID)
	var (
		err         error
		ctx         = context.Background()
		logr        = logx.WithContext(ctx)
		currentTime = int(time.Now().Unix())
	)
	// 判断是否开始
	if n.data.StartAt != 0 && n.data.StartAt > currentTime {
		logr.Infof("NotifyJob: %d not start", n.data.ID)
		return
	}
	// 判断是否过期
	if n.data.EndAt != 0 && n.data.EndAt < currentTime {
		logr.Infof("NotifyJob: %d is expired", n.data.ID)
		// 完成Job
		n.Complete()
		return
	}
	// 判断是否已经发送完成
	if n.data.MaxNotifyCount != -1 && n.data.NotifyCount >= n.data.MaxNotifyCount {
		logr.Infof("NotifyJob: [%d] has been sent, max[%d], current[%d]", n.data.ID, n.data.MaxNotifyCount, n.data.NotifyCount)
		// 完成Job
		n.Complete()
		return
	}

	logr.Infof("NotifyJob: [%d] start, max[%d], current[%d]", n.data.ID, n.data.MaxNotifyCount, n.data.NotifyCount)

	// 获取消息发送器
	sender := GetSender(n.channel.Sender)
	if sender == nil {
		// 通道不存在
		logr.Errorf("NotifyJob: [%d] sender [%s] not found", n.data.ID, n.channel.Sender)
		_ = n.done(n.data, errors.WithMessagef(types.ErrSenderNotFount, "sender: %s", n.channel.Sender))
		return
	}

	// 执行发送
	err = sender.Send(n.channel.Topic, n.channel.WechatUserID, n.data.Title, n.data.Content)
	if err != nil {
		// 发送失败
		logr.Errorf("NotifyJob: [%d] send failed, err: [%s]", n.data.ID, err.Error())
		_ = n.done(n.data, errors.WithMessage(types.ErrSendFailed, err.Error()))
		return
	}
	// 发送成功
	n.data.NotifyCount++
	n.data.LastNotifyAt = currentTime
	_ = n.done(n.data, err)

	// 判断是否已经发送完成
	if n.data.MaxNotifyCount != -1 && n.data.NotifyCount >= n.data.MaxNotifyCount {
		logr.Infof("NotifyJob: [%d] has been sent, max[%d], current[%d]", n.data.ID, n.data.MaxNotifyCount, n.data.NotifyCount)
		// 完成Job
		n.Complete()
		return
	}
}

func (n *NotifyJob) Complete() {
	n.completeJob(n.data)
}
