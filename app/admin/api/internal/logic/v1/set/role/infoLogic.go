package role

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.AdminRoleInfoReq) (resp *types.AdminRole, err error) {
	adminRole := model.AdminRole{}
	err = l.svcCtx.DB.WithContext(l.ctx).Where("id =?", req.Id).First(&adminRole).Error
	if err != nil {
		return nil, errors.Wrap(errx.NewCodeError(errx.Error), "信息不存在")
	}

	var permissions []model.AdminRolePermission
	l.svcCtx.DB.WithContext(l.ctx).Where("role_id=?", adminRole.ID).Find(&permissions)
	var per []int32
	for _, v := range permissions {
		per = append(per, v.MenuID)
	}
	adminRoleInfo := new(types.AdminRole)
	helper.ExchangeStruct(adminRole, &adminRoleInfo)
	adminRoleInfo.Permission = per
	return adminRoleInfo, nil
}
