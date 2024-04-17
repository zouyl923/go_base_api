package model

import (
	"time"
)

const TableNameArticle = "article"

// Article 文章表
type Article struct {
	Uuid      string    `gorm:"column:uuid;primaryKey;" json:"uuid"`
	Title     string    `gorm:"column:title;not null;comment:标题" json:"title"`           // 标题
	Cover     string    `gorm:"column:cover;comment:封面" json:"cover"`                    // 封面
	State     int32     `gorm:"column:state;comment:审核状态" json:"state"`                  // 审核状态
	ViewNum   int64     `gorm:"column:view_num;comment:浏览数" json:"view_num"`             // 浏览数
	LikeNum   int64     `gorm:"column:like_num;comment:点赞数" json:"like_num"`             // 点赞数
	IsHid     int32     `gorm:"column:is_hid;not null;comment:是否禁用：1是 0否" json:"is_hid"` // 是否禁用：1是 0否
	IsDel     int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"` // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int32     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间
	DetailInfo  AdminRole `gorm:"foreignKey:id;references:role_id" json:"role_info"`         // 角色信息
}

// TableName Admin table name
func (*Article) TableName() string {
	return TableNameArticle
}
