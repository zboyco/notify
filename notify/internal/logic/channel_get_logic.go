package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelGetLogic {
	return &ChannelGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelGetLogic) ChannelGet(req *types.ChannelByNameRequest) (resp *types.Channel, err error) {
	// todo: add your logic here and delete this line

	// 查询channel
	channel := &model.Channel{}
	channel.Name = req.ChannelName
	if err = channel.FetchByName(l.svcCtx.DB); err != nil {
		return nil, err
	}
	resp = &types.Channel{
		BaseModel: types.BaseModel{
			ID:        channel.ID,
			CreatedAt: channel.CreatedAt,
			UpdatedAt: channel.UpdatedAt,
		},
		Name:         channel.Name,
		Sender:       channel.Sender,
		WechatUserID: channel.WechatUserID,
		Topic:        channel.Topic,
		SubscribeURL: channel.SubscribeURL,
		SubscribeQr:  channel.SubscribeQr,
		Remark:       channel.Remark,
	}
	return
}
