package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/notify"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyCreateLogic {
	return &NotifyCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyCreateLogic) NotifyCreate(req *types.NotifyCreateRequest) error {
	// todo: add your logic here and delete this line
	notifyInfo := &model.Notify{
		WechatUserID: req.WechatUserID,
		TopicID:      req.TopicID,
		Title:        req.Title,
		Content:      req.Content,
		Loop:         req.Loop,
		StartAt:      req.StartAt,
		EndAt:        req.EndAt,
		Spec:         req.Spec,
		NotifyCount:  req.NotifyCount,
		LastNotifyAt: req.LastNotifyAt,
	}
	// 创建
	if err := notifyInfo.Create(l.svcCtx.DB); err != nil {
		return err
	}
	// 添加定时任务
	go func() {
		if req.Loop {
			err := l.svcCtx.CronJobRunner.AddJob(notifyInfo.ID, notifyInfo.Spec, notify.NewNotifyJob(l.svcCtx.DB, notifyInfo))
			if err != nil {
				l.Logger.Error(err)
			}
		}
	}()
	return nil
}
