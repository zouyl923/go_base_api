// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package commonservice

import (
	"context"

	"blog/app/article/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Article     = rpc.Article
	Category    = rpc.Category
	CategoryRes = rpc.CategoryRes
	DeleteReq   = rpc.DeleteReq
	Detail      = rpc.Detail
	EmptyReq    = rpc.EmptyReq
	EmptyRes    = rpc.EmptyRes
	ExamineReq  = rpc.ExamineReq
	InfoReq     = rpc.InfoReq
	PageData    = rpc.PageData
	SearchReq   = rpc.SearchReq
	UpdateReq   = rpc.UpdateReq
	UploadReq   = rpc.UploadReq
	UploadRes   = rpc.UploadRes
	User        = rpc.User

	CommonService interface {
		// 公共文章
		PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error)
		CategoryList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CategoryRes, error)
		Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error)
	}

	defaultCommonService struct {
		cli zrpc.Client
	}
)

func NewCommonService(cli zrpc.Client) CommonService {
	return &defaultCommonService{
		cli: cli,
	}
}

// 公共文章
func (m *defaultCommonService) PageList(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*PageData, error) {
	client := rpc.NewCommonServiceClient(m.cli.Conn())
	return client.PageList(ctx, in, opts...)
}

func (m *defaultCommonService) CategoryList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*CategoryRes, error) {
	client := rpc.NewCommonServiceClient(m.cli.Conn())
	return client.CategoryList(ctx, in, opts...)
}

func (m *defaultCommonService) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*Article, error) {
	client := rpc.NewCommonServiceClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}
