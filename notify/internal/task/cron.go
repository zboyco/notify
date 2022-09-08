package task

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/zboyco/notify/notify/internal/types"
	"gorm.io/gorm"
)

var (
	onceCronJobRunner sync.Once
	cronJobRunner     *CronJobRunner
)

// CronJobRunner is a cron job runner.
type CronJobRunner struct {
	cron        *cron.Cron
	jobEntryMap map[uint]cron.EntryID
	lock        *sync.RWMutex
}

func NewCronJobRunner(db *gorm.DB) *CronJobRunner {
	// 保证只有一个实例
	onceCronJobRunner.Do(func() {
		// 使用秒
		c := cron.New(cron.WithSeconds())
		c.Start()

		cronJobRunner = &CronJobRunner{
			cron:        c,
			jobEntryMap: make(map[uint]cron.EntryID),
			lock:        &sync.RWMutex{},
		}
	})
	return cronJobRunner
}

// 添加定时任务
func (c *CronJobRunner) AddJob(jobID uint, spec string, job cron.Job) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.jobEntryMap[jobID]; ok {
		return errors.WithMessage(types.ErrCronJob, "jobID already exists")
	}
	entryID, err := c.cron.AddJob(spec, job)
	if err != nil {
		return err
	}
	c.jobEntryMap[jobID] = entryID
	return nil
}

// 删除定时任务
func (c *CronJobRunner) RemoveJob(jobID uint) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if entryID, ok := c.jobEntryMap[jobID]; ok {
		c.cron.Remove(entryID)
		delete(c.jobEntryMap, jobID)
		return nil
	}
	return errors.WithMessage(types.ErrCronJob, "job not exists")
}

// 清空定时任务
func (c *CronJobRunner) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	// 停止定时任务
	_ = c.cron.Stop()
	// 初始化
	c.cron = cron.New(cron.WithSeconds())
	c.cron.Start()
	// 清空
	c.jobEntryMap = make(map[uint]cron.EntryID)
}
