package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"
	"github.com/zboyco/notify/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelListLogic {
	return &ChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelListLogic) ChannelList(req *types.ChannelListRequest) (resp *types.ChannelListResponse, err error) {
	// todo: add your logic here and delete this line

	// 查询channels
	channel := &model.Channel{}
	channels, err := channel.List(l.svcCtx.DB, utils.Pager{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}
	count, err := channel.Count(l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	// 返回结果
	resp = &types.ChannelListResponse{
		Total: count,
		Data:  make([]*types.Channel, 0, len(channels)),
	}
	for _, v := range channels {
		resp.Data = append(resp.Data, &types.Channel{
			BaseModel: types.BaseModel{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Name:         v.Name,
			Sender:       v.Sender,
			WechatUserID: v.WechatUserID,
			Topic:        v.Topic,
			SubscribeURL: v.SubscribeURL,
			SubscribeQr:  v.SubscribeQr,
			Remark:       v.Remark,
		})
	}

	return
}
