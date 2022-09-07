package notify

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zboyco/notify/notify/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type NotifyJob struct {
	db        *gorm.DB
	data      *model.Notify
	deleteJob func()
}

func NewNotifyJob(db *gorm.DB, data *model.Notify, deleteJobFunc func()) *NotifyJob {
	return &NotifyJob{
		db:        db,
		data:      data,
		deleteJob: deleteJobFunc,
	}
}

func (n *NotifyJob) Run() {
	ctx := context.Background()
	logr := logx.WithContext(ctx)
	currentTime := int(time.Now().Unix())
	// 判断是否开始
	if n.data.StartAt != 0 && n.data.StartAt > currentTime {
		logr.Infof("NotifyJob: %d not start", n.data.ID)
		return
	}
	// 判断是否过期
	if n.data.EndAt != 0 && n.data.EndAt < currentTime {
		logr.Infof("NotifyJob: %d is expired", n.data.ID)
		n.deleteJob()
		return
	}
	// 判断是否已经发送完成
	if n.data.MaxNotifyCount != 0 && n.data.NotifyCount >= n.data.MaxNotifyCount {
		logr.Infof("NotifyJob: %d has been sent, max[%d], current[%d]", n.data.ID, n.data.MaxNotifyCount, n.data.NotifyCount)
		n.deleteJob()
		return
	}

	logr.Info("run job notify id ", n.data.ID, " notify count ", n.data.NotifyCount)
	// 采用事务创建日志、更新通知次数、更新最后通知时间
	err := n.db.Transaction(func(tx *gorm.DB) error {
		// 获取消息发送器
		sender := GetSender(n.data.Channel)
		if sender == nil {
			return errors.New(fmt.Sprintf("sender %s not found", n.data.Channel))
		}
		// create notify log
		notifyLog := &model.NotifyLog{
			NotifyID: n.data.ID,
			Channel:  sender.Channel(),
			Status:   1,
			NotifyAt: currentTime,
		}
		// 发送通知
		err := sender.Send(n.data.Topic, n.data.WechatUserID, n.data.Title, n.data.Content)
		if err != nil {
			notifyLog.Log = err.Error()
			notifyLog.Status = 2
		}
		// 创建通知日志
		err = notifyLog.Create(tx)
		if err != nil {
			return err
		}
		// update notify count
		n.data.NotifyCount++
		// update notify last notify at
		n.data.LastNotifyAt = currentTime
		if err := n.data.Update(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logr.Error(err)
		return
	}
	// 判断是否已经发送完成
	if n.data.MaxNotifyCount != 0 && n.data.NotifyCount >= n.data.MaxNotifyCount {
		logr.Infof("NotifyJob: %d has been sent, max[%d], current[%d]", n.data.ID, n.data.MaxNotifyCount, n.data.NotifyCount)
		n.deleteJob()
		return
	}
}

func (n *NotifyJob) Delete() {
	n.deleteJob()
}
