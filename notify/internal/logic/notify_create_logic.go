package logic

import (
	"context"
	"time"

	"github.com/pkg/errors"
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

	// 判断结束时间是否大于当前时间
	if req.EndAt != 0 && req.EndAt < int(time.Now().Unix()) {
		return errors.WithMessage(types.ErrNotifyCreateFailed, "结束时间不能小于当前时间")
	}
	// 判断结束时间是否大于开始时间
	if req.EndAt != 0 && req.EndAt < req.StartAt {
		return errors.WithMessage(types.ErrNotifyCreateFailed, "结束时间不能小于开始时间")
	}

	// 查询channel是否存在
	channel := &model.Channel{}
	channel.Name = req.ChannelName
	if err := channel.FetchByName(l.svcCtx.DB); err != nil {
		return errors.WithMessage(types.ErrNotifyCreateFailed, err.Error())
	}

	notifyData := &model.Notify{
		ChannelID:      channel.ID,
		Title:          req.Title,
		Content:        req.Content,
		MaxNotifyCount: req.MaxNotifyCount,
		NotifyCount:    req.NotifyCount,
		StartAt:        req.StartAt,
		EndAt:          req.EndAt,
		Spec:           req.Spec,
		LastNotifyAt:   req.LastNotifyAt,
		Completed:      false,
	}
	// 创建
	if err := notifyData.Create(l.svcCtx.DB); err != nil {
		return err
	}
	// 添加定时任务
	AddNotifyToCron(l.ctx, l.svcCtx, notifyData)
	return nil
}
