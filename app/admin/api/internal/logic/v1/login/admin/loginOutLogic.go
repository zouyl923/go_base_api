package admin

import (
	"blog/app/verify/rpc/pb/rpc"
	"context"
	"net/http"

	"blog/app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginOutLogic) LoginOut(r *http.Request) error {
	adminId := r.Header.Get("admin_id")
	_, err := l.svcCtx.VerifyRpc.RemoveToken(l.ctx, &rpc.RemoveTokenReq{
		Server:   "admin",
		Platform: "all",
		Key:      adminId,
	})
	return err
}
