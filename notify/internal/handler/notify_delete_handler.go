package handler

import (
	"net/http"

	"github.com/zboyco/notify/notify/internal/logic"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func notifyDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyByIDRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewNotifyDeleteLogic(r.Context(), svcCtx)
		err := l.NotifyDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
