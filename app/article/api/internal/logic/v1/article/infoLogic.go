package article

import (
	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
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

func (l *InfoLogic) Info(req *types.ArticleInfoReq) (resp *types.Article, err error) {
	uuid := req.Uuid
	var info model.Article
	err = l.svcCtx.DB.WithContext(l.ctx).
		Where("uuid=?", uuid).
		Preload("DetailInfo").
		First(&info).Error
	if err != nil {
		return nil, errx.NewCodeError(errx.NotFundError)
	}
	cInfo := new(types.Article)
	helper.ExchangeStruct(info, cInfo)
	return cInfo, nil
}
