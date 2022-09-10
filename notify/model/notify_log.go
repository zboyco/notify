package model

import (
	"github.com/zboyco/notify/utils"
	"gorm.io/gorm"
)

type NotifyLog struct {
	BaseModel

	NotifyID  uint   `gorm:"not null;index"`           // 通知ID
	ChannelID uint   `gorm:"not null;index;default:0"` // 渠道ID
	Log       string `gorm:"not null;default:''"`      // 日志
	Status    int    `gorm:"not null;index;default:0"` // 通知状态
	NotifyAt  int    `gorm:"not null;default:0"`       // 通知时间
}

// 通过ID获取
func (t *NotifyLog) FetchByID(db *gorm.DB) error {
	return db.First(t, t.ID).Error
}

// 列表
func (t *NotifyLog) List(db *gorm.DB, pager utils.Pager) ([]NotifyLog, error) {
	var notifyLogs []NotifyLog
	err := db.Where(t).Limit(pager.Limit).Offset(pager.Offset).Find(&notifyLogs).Error
	if err != nil {
		return nil, err
	}
	return notifyLogs, nil
}

// 创建
func (t *NotifyLog) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

// 更新
func (t *NotifyLog) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

// 通过ID删除
func (t *NotifyLog) DeleteByID(db *gorm.DB) error {
	return db.Delete(t, t.ID).Error
}
