package svc

import (
	"blog/app/admin/api/internal/config"
	"blog/app/admin/api/internal/middleware"
	"blog/app/verify/rpc/client/verifyservice"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	AuthMiddleware rest.Middleware
	DB             *gorm.DB
	Cache          *redis.Redis
	VerifyRpc      verifyservice.VerifyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	verifyRpc := verifyservice.NewVerifyService(zrpc.MustNewClient(c.VerifyRpc))
	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		AuthMiddleware: middleware.NewAuthMiddleware(c, verifyRpc).Handle,
		DB:             GetOrm(c),
		Cache:          GetRedis(c),
		VerifyRpc:      verifyRpc,
	}
}

func GetRedis(c config.Config) *redis.Redis {
	rds := redis.MustNewRedis(c.Redis)
	return rds
}

func GetOrm(c config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(GetDsn(c)))
	if err != nil {
		//中断程序并报错
		panic(err)
	}
	db.Logger.LogMode(4)
	return db
}

func GetDsn(c config.Config) string {
	conf := c.MySql
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
	)
}
