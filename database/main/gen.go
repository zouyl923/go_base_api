package main

import (
	"blog/database"
	"blog/database/model"
	"gorm.io/gen"
)

var Dir = "../"

func main() {
	Migrate()
}

/**
 * 根据结构体生成表
 */
func Migrate() {
	db := database.NewDb()
	db.Migrator().AutoMigrate(
		&model.Admin{},
		&model.AdminLog{},
		&model.AdminMenu{},
		&model.AdminMessage{},
		&model.AdminPassword{},
		&model.AdminPermission{},
		&model.AdminRole{},
		&model.AdminRolePermission{},
		&model.Config{},
		&model.Article{},
		&model.ArticleDetail{},
		&model.ArticleCategory{},
		&model.User{},
	)
}

/**
 * 根据表生成model
 */
func Generator() {
	g := gen.NewGenerator(gen.Config{
		// 生成目录
		OutPath:      Dir + "query",
		ModelPkgPath: Dir + "model",
		// generate mode
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	db := database.NewDb()
	// reuse your gorm db
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	// 执行并生成代码
	g.Execute()
}
