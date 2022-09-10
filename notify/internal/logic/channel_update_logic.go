package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelUpdateLogic {
	return &ChannelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelUpdateLogic) ChannelUpdate(req *types.ChannelUpdateRequest) error {
	// todo: add your logic here and delete this line

	// 查询channel
	channel := &model.Channel{}
	channel.Name = req.ChannelName
	if err := channel.FetchByName(l.svcCtx.DB); err != nil {
		return err
	}
	// 更新channel
	channel.Sender = req.Sender
	channel.WechatUserID = req.WechatUserID
	channel.Topic = req.Topic
	channel.SubscribeURL = req.SubscribeURL
	channel.SubscribeQr = req.SubscribeQr
	channel.Remark = req.Remark
	return channel.Update(l.svcCtx.DB)
}
