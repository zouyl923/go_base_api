package middleware

import (
	"blog/app/admin/api/internal/config"
	"blog/app/verify/rpc/client/verifyservice"
	"blog/common/response"
	"blog/common/response/errx"
	"context"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Config    config.Config
	VerifyRpc verifyservice.VerifyService
}

func NewAuthMiddleware(c config.Config, rpc verifyservice.VerifyService) *AuthMiddleware {
	return &AuthMiddleware{
		Config:    c,
		VerifyRpc: rpc,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		authRes, err := m.VerifyRpc.Auth(context.Background(), &verifyservice.AuthReq{
			Server: "admin",
			Token:  token,
		})
		if err != nil {
			resp := response.Response{
				Code:    errx.LoginError,
				Message: errx.GetCnMessage(errx.LoginError),
			}
			str, _ := json.Marshal(resp)
			w.Write(str)
			return
		}
		//追加参数
		ctx := r.Context()
		ctx = context.WithValue(ctx, "adminId", authRes.Key)
		newR := r.WithContext(ctx)
		next(w, newR)
	}
}
