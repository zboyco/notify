package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyGetLogic {
	return &NotifyGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyGetLogic) NotifyGet(req *types.NotifyByIDRequest) (resp *types.Notify, err error) {
	// todo: add your logic here and delete this line
	// 查询notify
	notify := &model.Notify{}
	notify.ID = req.NotifyID
	if err = notify.FetchByID(l.svcCtx.DB); err != nil {
		return nil, err
	}
	resp = &types.Notify{
		BaseModel: types.BaseModel{
			ID:        notify.ID,
			CreatedAt: notify.CreatedAt,
			UpdatedAt: notify.UpdatedAt,
		},
		ChannelID:      notify.ChannelID,
		Title:          notify.Title,
		Content:        notify.Content,
		MaxNotifyCount: notify.MaxNotifyCount,
		NotifyCount:    notify.NotifyCount,
		StartAt:        notify.StartAt,
		EndAt:          notify.EndAt,
		Spec:           notify.Spec,
		LastNotifyAt:   notify.LastNotifyAt,
		Completed:      notify.Completed,
	}
	return
}
