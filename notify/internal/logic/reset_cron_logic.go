package logic

import (
	"context"

	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zboyco/notify/notify/model"
	"github.com/zboyco/notify/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetCronLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetCronLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetCronLogic {
	return &ResetCronLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetCronLogic) ResetCron(req *types.Auth) error {
	// todo: add your logic here and delete this line
	// 清空cron
	l.svcCtx.CronJobRunner.Clear()
	// 添加任务
	go func() {
		logr := l.Logger
		notifyModel := &model.Notify{}
		notifies, err := notifyModel.List(l.svcCtx.DB, utils.Pager{Limit: -1})
		if err != nil {
			logr.Errorf("init list notifies error: %v", err)
			return
		}
		for i := range notifies {
			AddNotifyToCron(l.ctx, l.svcCtx, &notifies[i])
		}
	}()
	return nil
}
