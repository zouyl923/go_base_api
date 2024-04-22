package category

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllListLogic {
	return &AllListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllListLogic) AllList() (resp []types.Category, err error) {
	var list []model.ArticleCategory
	l.svcCtx.DB.WithContext(l.ctx).
		Where("is_del", 0).
		Where("is_hid", 0).
		Find(&list)
	var cList []types.Category
	helper.ExchangeStruct(list, &cList)
	return cList, nil
}
