package svc

import (
	"blog/app/article/api/internal/config"
	"blog/app/article/api/internal/middleware"
	articleRpcClient "blog/app/article/rpc/rpcClient"
	verifyRpcClient "blog/app/verify/rpc/rpcClient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	CorsMiddleware      rest.Middleware
	UserAuthMiddleware  rest.Middleware
	AdminAuthMiddleware rest.Middleware
	ArticleRpc          articleRpcClient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	articleRpc := articleRpcClient.NewRpc(zrpc.MustNewClient(c.ArticleRpc))
	verifyRpc := verifyRpcClient.NewRpc(zrpc.MustNewClient(c.VerifyRpc))
	return &ServiceContext{
		Config:             c,
		CorsMiddleware:     middleware.NewCorsMiddleware().Handle,
		UserAuthMiddleware: middleware.NewUserAuthMiddleware(c, verifyRpc).Handle,
		ArticleRpc:         articleRpc,
	}
}
