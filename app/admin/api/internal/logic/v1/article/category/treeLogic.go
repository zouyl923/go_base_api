package category

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeLogic {
	return &TreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeLogic) Tree() (resp []types.ArticleCategroy, err error) {
	// todo: add your logic here and delete this line

	return
}
