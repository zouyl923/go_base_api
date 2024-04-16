package menu

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

func (l *InfoLogic) Info(req *types.MenuInfoReq) (resp *types.Menu, err error) {
	info := model.AdminMenu{}
	err = l.svcCtx.DB.WithContext(l.ctx).Where("id = ?", req.Id).First(&info).Error
	if err != nil {
		return nil, errors.Wrap(errx.NewCodeError(errx.Error), "信息不存在！")
	}

	cInfo := new(types.Menu)
	helper.ChangeToStruct(info, &cInfo)
	return cInfo, nil
}
