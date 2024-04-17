package model

import (
	"time"
)

const TableNameAdmin = "admin"

// Admin 管理员表
type Admin struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null;comment:账号名" json:"name"`            // 账号名
	Phone     string    `gorm:"column:phone;not null;comment:手机号" json:"phone"`          // 手机号
	Password  string    `gorm:"column:password;not null;comment:密码" json:"password"`     // 密码
	RoleID    int32     `gorm:"column:role_id;not null;comment:密码" json:"role_id"`       // 密码
	IsHid     int32     `gorm:"column:is_hid;not null;comment:是否禁用：1是 0否" json:"is_hid"` // 是否禁用：1是 0否
	IsDel     int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"` // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int32     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间

	RoleInfo AdminRole `gorm:"foreignKey:id;references:role_id" json:"role_info"` // 角色信息
}

// TableName Admin table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
