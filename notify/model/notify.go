package model

import (
	"github.com/zboyco/notify/utils"
	"gorm.io/gorm"
)

type Notify struct {
	BaseModel

	ChannelID      uint   `gorm:"not null;index"`               // 通知渠道
	Title          string `gorm:"not null;default:''"`          // 标题
	Content        string `gorm:"not null;default:''"`          // 内容
	MaxNotifyCount int    `gorm:"not null;default:1"`           // 最大通知次数，0为不限制
	NotifyCount    int    `gorm:"not null;default:0"`           // 已通知次数
	StartAt        int    `gorm:"not null;default:0"`           // 开始时间
	EndAt          int    `gorm:"not null;default:0"`           // 结束时间
	Spec           string `gorm:"not null;default:''"`          // Cron表达式
	LastNotifyAt   int    `gorm:"not null;default:0"`           // 最后通知时间
	Completed      bool   `gorm:"not null;index;default:false"` // 完成标识
}

// 通过ID获取
func (t *Notify) FetchByID(db *gorm.DB) error {
	return db.First(t, t.ID).Error
}

// 统计
func (t *Notify) Count(db *gorm.DB, zeroFields ...interface{}) (int64, error) {
	var count int64
	err := db.Model(t).Where(t, zeroFields...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 列表
func (t *Notify) List(db *gorm.DB, pager utils.Pager, orderFields interface{}, zeroFields ...interface{}) ([]Notify, error) {
	var notifies []Notify
	err := db.Where(t, zeroFields...).Order(orderFields).Limit(pager.Limit).Offset(pager.Offset).Find(&notifies).Error
	if err != nil {
		return nil, err
	}
	return notifies, nil
}

// 列表(自定义条件)
func (t *Notify) ListByQuery(db *gorm.DB, pager utils.Pager, where interface{}, args ...interface{}) ([]Notify, error) {
	var notifies []Notify
	err := db.Where(where, args...).Limit(pager.Limit).Offset(pager.Offset).Find(&notifies).Error
	if err != nil {
		return nil, err
	}
	return notifies, nil
}

// 创建
func (t *Notify) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

// 更新
func (t *Notify) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

// 通过ID删除
func (t *Notify) DeleteByID(db *gorm.DB) error {
	return db.Delete(t, t.ID).Error
}
