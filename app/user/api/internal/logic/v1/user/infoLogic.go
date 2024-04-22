package user

import (
	"blog/app/user/rpc/pb/rpc"
	"blog/common/helper"
	"context"
	"net/http"
	"strconv"

	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"

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

func (l *InfoLogic) Info(r *http.Request) (resp *types.User, err error) {
	userId := r.Header.Get("user_id")
	uId, err := strconv.ParseInt(userId, 10, 64)
	info, err := l.svcCtx.UserRpc.Info(l.ctx, &rpc.InfoReq{
		UserId: uId,
	})
	if err != nil {
		return nil, err
	}
	cInfo := types.User{}
	helper.ExchangeStruct(info, &cInfo)
	return &cInfo, nil
}
