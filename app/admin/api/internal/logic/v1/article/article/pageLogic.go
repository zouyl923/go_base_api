package article

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.ArticleSearchReq) (resp *types.ArticlePageList, err error) {
	// todo: add your logic here and delete this line

	return
}
