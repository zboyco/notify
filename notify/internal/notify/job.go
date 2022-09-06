package notify

import (
	"context"
	"time"

	"github.com/zboyco/notify/notify/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type NotifyJob struct {
	db   *gorm.DB
	data *model.Notify
}

func NewNotifyJob(db *gorm.DB, data *model.Notify) *NotifyJob {
	return &NotifyJob{
		db:   db,
		data: data,
	}
}

func (n *NotifyJob) Run() {
	ctx := context.Background()
	logr := logx.WithContext(ctx)
	currentTime := int(time.Now().Unix())
	// 判断是否在有效期内
	if (n.data.StartAt != 0 && n.data.StartAt > currentTime) ||
		(n.data.EndAt != 0 && n.data.EndAt < currentTime) {
		logr.Infof("NotifyJob: %d is not in the effective period", n.data.ID)
		return
	}

	logr.Info("run job notify id ", n.data.ID, " notify count ", n.data.NotifyCount)
	// 采用事务创建日志、更新通知次数、更新最后通知时间
	err := n.db.Transaction(func(tx *gorm.DB) error {
		// 循环消息发送器
		for _, sender := range Senders {
			// create notify log
			notifyLog := &model.NotifyLog{
				NotifyID: n.data.ID,
				Channel:  sender.Channel(),
				Status:   1,
				NotifyAt: currentTime,
			}
			// 发送通知
			err := sender.Send(n.data.TopicID, n.data.WechatUserID, n.data.Title, n.data.Content)
			if err != nil {
				notifyLog.Log = err.Error()
				notifyLog.Status = 2
			}
			// 创建通知日志
			err = notifyLog.Create(tx)
			if err != nil {
				return err
			}
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
	}
}
