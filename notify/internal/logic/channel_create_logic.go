package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelCreateLogic {
	return &ChannelCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelCreateLogic) ChannelCreate(req *types.ChannelCreateRequest) error {
	// todo: add your logic here and delete this line

	// 检查接收者
	if req.WechatUserID == "" && req.Topic == "" {
		return errors.WithMessage(types.ErrNotifyCreateFailed, "接收者不能为空")
	}

	// 创建channel
	channel := &model.Channel{
		Name:         req.Name,
		Sender:       req.Sender,
		WechatUserID: req.WechatUserID,
		Topic:        req.Topic,
		SubscribeURL: req.SubscribeURL,
		SubscribeQr:  req.SubscribeQr,
		Remark:       req.Remark,
	}
	return channel.Create(l.svcCtx.DB)
}
