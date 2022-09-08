package logic

import (
	"context"
	"time"

	"github.com/zboyco/notify/notify/internal/notify"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/task"
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
			return
		}
		// 如果通知时间小于当前时间，则跳过
		if notifyData.EndAt != 0 && notifyData.EndAt < currentTime {
			logr.Info("notify end time less than current time")
			return
		}
		// 创建job
		notifyJob := notify.NewNotifyJob(notifyData, deleteJobFunc(ctx, svcCtx.CronJobRunner), updateNotifyAndLog(svcCtx.DB))
		// 添加定时任务
		if err := svcCtx.CronJobRunner.AddJob(notifyData.ID, notifyData.Spec, notifyJob); err != nil {
			logr.Errorf("init add job %v error: %v", notifyData.ID, err)
		}
	}()
}

// 删除cron中的notify
func deleteJobFunc(ctx context.Context, runner *task.CronJobRunner) func(*model.Notify) {
	logr := logx.WithContext(ctx)
	return func(notifyData *model.Notify) {
		logr.Infof("notify job %d done", notifyData.ID)
		if err := runner.RemoveJob(notifyData.ID); err != nil {
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
				NotifyID: notifyData.ID,
				Channel:  notifyData.Channel,
				Status:   1,
				NotifyAt: currentTime,
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
