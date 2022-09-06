package handler

import (
	"net/http"

	"github.com/zboyco/notify/notify/internal/logic"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func wxpusherCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WxPusherCallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWxpusherCallbackLogic(r.Context(), svcCtx)
		err := l.WxpusherCallback(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
