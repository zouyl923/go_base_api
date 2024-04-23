package user

import (
	"blog/app/article/rpc/pb/rpc"
	"blog/common/response/errx"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) error {
	userId := l.ctx.Value("userId").(int64)
	_, err := l.svcCtx.ArticleUserRpc.Delete(l.ctx, &rpc.DeleteReq{
		Uuid:   req.Uuid,
		UserId: userId,
	})
	if err != nil {
		return errx.NewMessageError(err.Error())
	}
	return nil
}
