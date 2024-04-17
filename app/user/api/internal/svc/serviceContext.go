package svc

import (
	"blog/app/user/api/internal/config"
	"blog/app/user/api/internal/middleware"
	"blog/app/user/rpc/rpcClient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	UserRpc        rpcClient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		UserRpc:        rpcClient.NewRpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
