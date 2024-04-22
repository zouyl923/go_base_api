package article

import (
	"blog/common/helper"
	"blog/database/model"
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

func (l *PageListLogic) PageList(req *types.ArticleSearchReq) (resp *types.ArticlePageList, err error) {
	page := req.Page
	pageSize := 10

	offset := (page - 1) * pageSize
	var list []model.Article
	var total int64

	model := l.svcCtx.DB.WithContext(l.ctx)
	if len(req.Keyword) > 0 {
		model = model.Where("title like ?", "%"+req.Keyword+"%")
	}
	model.Where("is_del = ?", 0).
		Where("is_hid = ?", 0).
		Offset(offset).
		Limit(pageSize).
		Preload("CategoryInfo").
		Find(&list).
		Count(&total)
	//数据格式转换
	var cList []types.Article
	helper.ExchangeStruct(list, &cList)
	resp = new(types.ArticlePageList)
	resp.Page = page
	resp.PageSize = pageSize
	resp.Total = total
	resp.Data = cList
	return resp, nil
}
