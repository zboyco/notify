package crontab

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/zboyco/notify/notify/internal/notify"
	"github.com/zboyco/notify/notify/model"
	"github.com/zboyco/notify/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	runner *CronJobRunner
)

type CronJobRunner struct {
	cron        *cron.Cron
	jobEntryMap map[uint]cron.EntryID
}

func NewCronJobRunner(db *gorm.DB) *CronJobRunner {
	// 保证只有一个实例
	once.Do(func() {
		// 使用秒
		c := cron.New(cron.WithSeconds())
		c.Start()

		runner = &CronJobRunner{
			cron:        c,
			jobEntryMap: make(map[uint]cron.EntryID),
		}
		// 初始化添加定时任务
		go func() {
			logr := logx.WithContext(context.Background())
			notifyModel := &model.Notify{
				Loop: true,
			}
			notifies, err := notifyModel.List(db, utils.Pager{Limit: -1})
			if err != nil {
				logr.Errorf("init list notifies error: %v", err)
				return
			}
			// 当前时间
			currentTime := int(time.Now().Unix())
			for i := range notifies {
				go func(notifyData *model.Notify) {
					// 如果通知时间小于当前时间，则跳过
					if notifyData.EndAt > currentTime {
						// 添加定时任务
						if err := runner.AddJob(notifyData.ID, notifyData.Spec, notify.NewNotifyJob(db, notifyData)); err != nil {
							logr.Errorf("init add job %v error: %v", notifyData.ID, err)
						}
					}
				}(&notifies[i])
			}
		}()
	})
	return runner
}

// 添加定时任务
func (c *CronJobRunner) AddJob(jobID uint, spec string, job cron.Job) error {
	if _, ok := c.jobEntryMap[jobID]; ok {
		return errors.New("jobID already exists")
	}
	entryID, err := c.cron.AddJob(spec, job)
	if err != nil {
		return err
	}
	c.jobEntryMap[jobID] = entryID
	return nil
}

// 删除定时任务
func (c *CronJobRunner) RemoveJob(notifyID uint) error {
	if entryID, ok := c.jobEntryMap[notifyID]; ok {
		c.cron.Remove(entryID)
		delete(c.jobEntryMap, notifyID)
		return nil
	}
	return errors.New("job not exists")
}
