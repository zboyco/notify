package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxpusherCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxpusherCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxpusherCallbackLogic {
	return &WxpusherCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxpusherCallbackLogic) WxpusherCallback(req *types.WxPusherCallbackRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
