// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zboyco/notify/notify/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v0/wxpusher/callback",
				Handler: wxpusherCallbackHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v0/notifies",
					Handler: notifyCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v0/notifies",
					Handler: notifyListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v0/notifies/:notifyID",
					Handler: notifyGetHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/v0/notifies/:notifyID",
					Handler: notifyDeleteHandler(serverCtx),
				},
			}...,
		),
	)
}
