package category

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ArticleCategroyUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
