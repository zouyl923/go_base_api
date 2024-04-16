package menu

import (
	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"github.com/pkg/errors"
	"time"

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

func (l *UpdateLogic) Update(req *types.MenuUpdateReq) error {
	data := model.AdminMenu{
		ID:        req.Id,
		ParentID:  req.ParentId,
		Name:      req.Name,
		URI:       req.Uri,
		Icon:      req.Icon,
		Weight:    req.Weight,
		CreatedAt: time.Now(),
	}
	err := l.svcCtx.DB.WithContext(l.ctx).Save(&data).Error
	if err != nil {
		return errors.Wrap(errx.NewCodeError(errx.Error), "操作失败！")
	}
	return nil
}
