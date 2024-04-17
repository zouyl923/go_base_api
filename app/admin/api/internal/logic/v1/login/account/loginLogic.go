package account

import (
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	resp = new(types.LoginRes)
	adminInfo := &model.Admin{}
	err = l.svcCtx.DB.WithContext(l.ctx).
		Where("name = ?", req.Username).
		Preload("RoleInfo").
		First(&adminInfo).Error
	if err != nil {
		return nil, errors.Wrap(errx.NewCodeError(errx.AdminNotFound), "账户不存在！")
	}

	check := helper.PasswordVerify(req.Password, adminInfo.Password)
	if check != true {
		return nil, errors.Wrap(errx.NewCodeError(errx.AdminNotFound), "密码错误！")
	}
	key := fmt.Sprintf("%d", adminInfo.ID)
	token, err := helper.GenToken(key, 2*60*60)
	if err != nil {
		return nil, errors.Wrap(errx.NewCodeError(errx.AdminNotFound), "token生成失败！")
	}

	cacheKey := "admin_token:" + key
	l.svcCtx.Cache.Setex(cacheKey, token, 30*24*60*60)

	admin := types.AdminInfo{}
	helper.ChangeToStruct(adminInfo, &admin)
	resp.AdminInfo = admin
	resp.Token = token
	return resp, nil
}
