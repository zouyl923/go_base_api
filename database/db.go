package database

import (
	"blog/common/file"
	"encoding/json"
	"fmt"
	"path"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Prefix   string `yaml:"prefix"`
	Charset  string `yaml:"charset"`
}

func NewDb() *gorm.DB {
	dsn := GetDsn()
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//禁用生成外键关联
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if Err != nil {
		panic(Err)
	}
	return DB
}

func GetConfig() Config {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(path.Dir(filename))
	confDir := dir + "/database/mysql.yaml"

	c := Config{}
	mConf, _ := file.LoadConfig(confDir, &c)
	//map 转 struct
	jConf, _ := json.Marshal(mConf)
	_ = json.Unmarshal(jConf, &c)
	return c
}

func GetDsn() string {
	c := GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Charset,
	)
}
