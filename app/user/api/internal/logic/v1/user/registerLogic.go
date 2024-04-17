package user

import (
	"blog/app/user/rpc/rpcClient"
	"blog/common/helper"
	"blog/common/response/errx"
	"context"
	"github.com/pkg/errors"

	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.LoginRes, err error) {
	info, err := l.svcCtx.UserRpc.Register(l.ctx, &rpcClient.RegisterReq{
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrap(errx.NewMessageError(err.Error()), "")
	}
	res := new(types.LoginRes)
	helper.ChangeToStruct(info, res)
	return res, nil
}
