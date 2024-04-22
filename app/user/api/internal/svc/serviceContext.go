package svc

import (
	"blog/app/user/api/internal/config"
	"blog/app/user/api/internal/middleware"
	userRpcClient "blog/app/user/rpc/rpcClient"
	verifyRpcClient "blog/app/verify/rpc/rpcClient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	AuthMiddleware rest.Middleware
	UserRpc        userRpcClient.Rpc
	VerifyRpc      verifyRpcClient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := userRpcClient.NewRpc(zrpc.MustNewClient(c.UserRpc))
	verifyRpc := verifyRpcClient.NewRpc(zrpc.MustNewClient(c.VerifyRpc))
	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		AuthMiddleware: middleware.NewAuthMiddleware(c, verifyRpc).Handle,
		UserRpc:        userRpc,
	}
}
