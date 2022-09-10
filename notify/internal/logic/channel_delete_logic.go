package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelDeleteLogic {
	return &ChannelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelDeleteLogic) ChannelDelete(req *types.ChannelByNameRequest) error {
	// todo: add your logic here and delete this line

	// 查询channel
	channel := &model.Channel{}
	channel.Name = req.ChannelName
	if err := channel.FetchByName(l.svcCtx.DB); err != nil {
		return err
	}
	// 删除channel
	return channel.DeleteByID(l.svcCtx.DB)
}
