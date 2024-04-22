package svc

import (
	"blog/app/user/rpc/internal/config"
	"blog/app/verify/rpc/rpcClient"
	"blog/common/helper"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Cache  *redis.Redis

	VerifyRpc rpcClient.Rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB:        GetOrm(c),
		Cache:     GetRedis(c),
		VerifyRpc: rpcClient.NewRpc(zrpc.MustNewClient(c.VerifyRpc)),
	}
}

func GetRedis(c config.Config) *redis.Redis {
	var conf redis.RedisConf
	helper.ExchangeStruct(c.Redis, &conf)
	rds := redis.MustNewRedis(conf)
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
