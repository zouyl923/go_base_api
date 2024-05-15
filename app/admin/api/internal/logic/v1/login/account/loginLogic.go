package account

import (
	"blog/app/verify/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"
	"strconv"

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
	if err != nil || adminInfo.ID < 1 {
		return nil, errors.Wrap(errx.NewCodeError(errx.AdminNotFound), "账户不存在！")
	}

	check := helper.PasswordVerify(req.Password, adminInfo.Password)
	if check != true {
		return nil, errx.NewCodeError(errx.AdminNotFound)
	}

	adminId := strconv.FormatInt(adminInfo.ID, 10)
	tokenRes, err := l.svcCtx.VerifyRpc.GenToken(l.ctx, &rpc.GenTokenReq{
		Server:   "admin",
		Platform: "all",
		Key:      adminId,
		Expire:   2 * 60 * 60,
	})
	if err != nil {
		return nil, err
	}

	admin := types.AdminInfo{}
	helper.ExchangeStruct(adminInfo, &admin)
	resp.AdminInfo = admin
	resp.Token = tokenRes.Token
	resp.RefreshToken = tokenRes.RefreshToken
	return resp, nil
}
