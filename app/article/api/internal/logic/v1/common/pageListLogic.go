package common

import (
	"blog/app/article/rpc/pb/rpc"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageListLogic {
	return &PageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageListLogic) PageList(req *types.SearchReq) (resp *types.PageList, err error) {
	l.svcCtx.ArticleRpc.PageList(l.ctx, &rpc.SearchReq{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Keyword:    req.Keyword,
		CategoryId: req.CategoryId,
	})

	return
}
