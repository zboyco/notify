package types

import "errors"

var (
	ErrNotifyCreateFailed = errors.New("通知创建失败")
	ErrCronJob            = errors.New("定时任务异常")
	ErrSenderNotFount     = errors.New("发送者不存在")
	ErrSendFailed         = errors.New("发送失败")
)
