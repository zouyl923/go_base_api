// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	v1login "blog/app/user/api/internal/handler/v1/login"
	v1user "blog/app/user/api/internal/handler/v1/user"
	"blog/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/login",
					Handler: v1login.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/register",
					Handler: v1login.RegisterHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/info",
					Handler: v1user.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/main",
					Handler: v1user.MainHandler(serverCtx),
				},
			}...,
		),
	)
}
