package logic

import (
	"context"
	"errors"
	"time"

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
	// 检查接收者
	if req.WechatUserID == "" && req.Topic == "" {
		return errors.New("接收者不能为空")
	}
	// 判断结束时间是否大于当前时间
	if req.EndAt != 0 && req.EndAt < int(time.Now().Unix()) {
		return errors.New("结束时间不能小于当前时间")
	}
	// 判断结束时间是否大于开始时间
	if req.EndAt != 0 && req.EndAt < req.StartAt {
		return errors.New("结束时间不能小于开始时间")
	}
	notifyData := &model.Notify{
		Channel:        req.Channel,
		WechatUserID:   req.WechatUserID,
		Topic:          req.Topic,
		Title:          req.Title,
		Content:        req.Content,
		MaxNotifyCount: req.MaxNotifyCount,
		NotifyCount:    req.NotifyCount,
		StartAt:        req.StartAt,
		EndAt:          req.EndAt,
		Spec:           req.Spec,
		LastNotifyAt:   req.LastNotifyAt,
	}
	// 创建
	if err := notifyData.Create(l.svcCtx.DB); err != nil {
		return err
	}
	// 添加定时任务
	go func() {
		notifyJob := notify.NewNotifyJob(l.svcCtx.DB, notifyData, func() {
			l.Logger.Infof("notify job %d done", notifyData.ID)
			if err := l.svcCtx.CronJobRunner.RemoveJob(notifyData.ID); err != nil {
				l.Logger.Errorf("remove job %d error: %v", notifyData.ID, err)
			}
		})
		err := l.svcCtx.CronJobRunner.AddJob(notifyData.ID, notifyData.Spec, notifyJob)
		if err != nil {
			l.Logger.Error(err)
		}
	}()
	return nil
}
