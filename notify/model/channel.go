package model

import (
	"github.com/zboyco/notify/utils"
	"gorm.io/gorm"
)

type Channel struct {
	BaseModel

	Name         string `gorm:"not null;uniqueIndex;default:''"` // 渠道名称
	Sender       string `gorm:"not null;default:''"`             // 发送者
	WechatUserID string `gorm:"not null;index;default:''"`       // 微信用户ID
	Topic        string `gorm:"not null;index;default:''"`       // 主题
	SubscribeURL string `gorm:"not null;default:''"`             // 订阅地址
	SubscribeQr  string `gorm:"not null;default:''"`             // 订阅二维码
	Remark       string `gorm:"not null;default:''"`             // 备注
}

// 通过ID获取
func (t *Channel) FetchByID(db *gorm.DB) error {
	return db.First(t, t.ID).Error
}

// 通过名称获取
func (t *Channel) FetchByName(db *gorm.DB) error {
	return db.First(t).Error
}

// 统计
func (t *Channel) Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(t).Where(t).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 列表
func (t *Channel) List(db *gorm.DB, pager utils.Pager) ([]Channel, error) {
	var channels []Channel
	err := db.Where(t).Limit(pager.Limit).Offset(pager.Offset).Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

// 创建
func (t *Channel) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

// 更新
func (t *Channel) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

// 通过ID删除
func (t *Channel) DeleteByID(db *gorm.DB) error {
	return db.Delete(t, t.ID).Error
}

// 通过名称删除
func (t *Channel) DeleteByName(db *gorm.DB) error {
	return db.Where("Name = ?", t.Name).Delete(t).Error
}
