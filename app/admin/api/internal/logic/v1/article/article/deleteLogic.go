package article

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"context"
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

func (l *DeleteLogic) Delete(req *types.ArticleDeleteReq) error {
	// todo: add your logic here and delete this line

	return nil
}
