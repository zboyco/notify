package logic

import (
	"context"
	"time"

	"github.com/zboyco/notify/notify/internal/notify"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 添加notify到cron
func AddNotifyToCron(ctx context.Context, svcCtx *svc.ServiceContext, notifyData *model.Notify) {
	// 当前时间
	currentTime := int(time.Now().Unix())
	logr := logx.WithContext(ctx)
	go func() {
		// 如果已通知次数已达最大通知次数，则跳过
		if notifyData.MaxNotifyCount != 0 && notifyData.NotifyCount >= notifyData.MaxNotifyCount {
			logr.Infof("notify %d max notify count reached", notifyData.ID)
			// 标识完成任务
			completeJobFunc(ctx, svcCtx)(notifyData)
			return
		}
		// 如果通知时间小于当前时间，则跳过
		if notifyData.EndAt != 0 && notifyData.EndAt < currentTime {
			logr.Info("notify end time less than current time")
			// 标识完成任务
			completeJobFunc(ctx, svcCtx)(notifyData)
			return
		}
		// 查询通知渠道
		channel := &model.Channel{}
		channel.ID = notifyData.ChannelID
		if err := channel.FetchByID(svcCtx.DB); err != nil {
			logr.Errorf("fetch channel %d error: %v", notifyData.ChannelID, err)
			return
		}
		// 创建job
		notifyJob := notify.NewNotifyJob(notifyData, channel, completeJobFunc(ctx, svcCtx), updateNotifyAndLog(svcCtx.DB))
		// 添加定时任务
		if err := svcCtx.CronJobRunner.AddJob(notifyData.ID, notifyData.Spec, notifyJob); err != nil {
			logr.Errorf("init add job %v error: %v", notifyData.ID, err)
		}
	}()
}

// 完成任务回调函数
func completeJobFunc(ctx context.Context, svcCtx *svc.ServiceContext) func(*model.Notify) {
	logr := logx.WithContext(ctx)
	return func(notifyData *model.Notify) {
		logr.Infof("notify job %d done", notifyData.ID)
		// 标识任务已完成
		notifyData.Completed = true
		if err := notifyData.Update(svcCtx.DB); err != nil {
			logr.Errorf("update notify %d completed error: %v", notifyData.ID, err)
		}
		// 删除cron中的notify
		if err := svcCtx.CronJobRunner.RemoveJob(notifyData.ID); err != nil {
			logr.Errorf("remove job %d error: %v", notifyData.ID, err)
		}
	}
}

// 更新notify信息并创建notify日志
func updateNotifyAndLog(db *gorm.DB) func(*model.Notify, error) error {
	return func(notifyData *model.Notify, err error) error {
		var currentTime = int(time.Now().Unix())
		// 采用事务创建日志、更新通知次数、更新最后通知时间
		return db.Transaction(func(tx *gorm.DB) error {
			// create notify log
			notifyLog := &model.NotifyLog{
				NotifyID:  notifyData.ID,
				ChannelID: notifyData.ChannelID,
				Status:    1,
				NotifyAt:  currentTime,
			}

			if err != nil {
				notifyLog.Log = err.Error()
				notifyLog.Status = 2
			}
			// 创建通知日志
			errLog := notifyLog.Create(tx)
			if errLog != nil {
				return errLog
			}
			// 发送失败，不更新通知次数
			if err != nil {
				return nil
			}

			// update notify
			return notifyData.Update(tx)
		})
	}
}
