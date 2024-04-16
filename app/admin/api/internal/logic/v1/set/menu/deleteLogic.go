package menu

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"

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

func (l *DeleteLogic) Delete(req *types.MenuDeleteReq) error {
	ids := make([]int64, 0)
	if req.Id > 0 {
		ids = append(ids, req.Id)
	}
	if len(req.Ids) > 0 {
		ids = req.Ids
	}
	var info []model.AdminMenu
	err := l.svcCtx.DB.WithContext(l.ctx).
		Where("id in (?)", ids).
		Find(&info).
		Update("is_del", 1).Error
	if err != nil {
		return errors.Wrap(errx.NewCodeError(errx.Error), "操作失败！")
	}
	return nil
}
