package admin

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"fmt"
	"net/http"

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

func (l *InfoLogic) Info(r *http.Request) (resp *types.AdminInfo, err error) {
	adminId := r.Header.Get("admin_id")
	fmt.Println(adminId)
	admin := &model.Admin{}
	err = l.svcCtx.DB.WithContext(l.ctx).
		Preload("RoleInfo").
		Where("id =?", adminId).
		First(admin).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.AdminNotFound)
	}

	adminInfo := new(types.AdminInfo)
	helper.ExchangeStruct(admin, &adminInfo)
	return adminInfo, nil
}
