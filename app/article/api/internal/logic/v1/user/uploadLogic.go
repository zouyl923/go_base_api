package user

import (
	"blog/app/article/api/internal/svc"
	"blog/app/article/api/internal/types"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(r *http.Request) (resp *types.UploadRes, err error) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//创建上传目录
	os.Mkdir("./upload", os.ModePerm)
	//创建上传文件
	f, err := os.Create("./upload/" + handler.Filename)
	defer f.Close()
	io.Copy(f, file)
	return &types.UploadRes{File: handler.Filename}, nil
}
