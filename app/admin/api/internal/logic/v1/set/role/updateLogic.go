package role

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.AdminRoleUpdateReq) error {
	role := model.AdminRole{
		ID:        req.Id,
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	tx := l.svcCtx.DB.WithContext(l.ctx).Begin()
	err := tx.Save(&role).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(errx.NewCodeError(errx.Error), "操作失败！")
	}
	var permission model.AdminRolePermission
	var rolePermission []model.AdminRolePermission
	for _, v := range req.Permission {
		permission.RoleID = int32(role.ID)
		permission.MenuID = int32(v)
		rolePermission = append(rolePermission, permission)
	}
	err = tx.Where("role_id=?", role.ID).Delete(&rolePermission).Create(&rolePermission).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(errx.NewCodeError(errx.Error), "操作失败！")
	}
	tx.Commit()
	return nil
}
