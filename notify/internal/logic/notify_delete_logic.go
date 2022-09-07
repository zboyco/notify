package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyDeleteLogic {
	return &NotifyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyDeleteLogic) NotifyDelete(req *types.NotifyByIDRequest) error {
	// todo: add your logic here and delete this line
	// 查询notify
	notify := &model.Notify{}
	notify.ID = req.NotifyID
	if err := notify.FetchByID(l.svcCtx.DB); err != nil {
		return err
	}
	// 移除定时任务
	err := l.svcCtx.CronJobRunner.RemoveJob(notify.ID)
	if err != nil {
		l.Logger.Error(err)
	}
	// 删除数据
	if err := notify.DeleteByID(l.svcCtx.DB); err != nil {
		return err
	}
	return nil
}
