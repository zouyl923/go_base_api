package category

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"context"

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

func (l *InfoLogic) Info(req *types.ArticleCategroyInfoReq) (resp *types.ArticleCategroy, err error) {
	// todo: add your logic here and delete this line

	return
}
