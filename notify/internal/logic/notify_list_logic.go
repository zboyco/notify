package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"
	"github.com/zboyco/notify/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyListLogic {
	return &NotifyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyListLogic) NotifyList(req *types.NotifyListRequest) (resp *types.NotifyListResponse, err error) {
	// todo: add your logic here and delete this line
	// 查询notify list
	notify := &model.Notify{}
	notifyList, err := notify.List(l.svcCtx.DB, utils.Pager{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}
	// 查询notify count
	count, err := notify.Count(l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	// 组装返回数据
	list := make([]*types.Notify, 0, len(notifyList))
	for _, v := range notifyList {
		list = append(list, &types.Notify{
			ID:             v.ID,
			Channel:        v.Channel,
			WechatUserID:   v.WechatUserID,
			Topic:          v.Topic,
			Title:          v.Title,
			Content:        v.Content,
			MaxNotifyCount: v.MaxNotifyCount,
			NotifyCount:    v.NotifyCount,
			StartAt:        v.StartAt,
			EndAt:          v.EndAt,
			Spec:           v.Spec,
			LastNotifyAt:   v.LastNotifyAt,
		})
	}
	resp = &types.NotifyListResponse{
		Total: count,
		Data:  list,
	}
	return
}
