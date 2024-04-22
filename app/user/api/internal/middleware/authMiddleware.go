package middleware

import (
	"blog/app/user/api/internal/config"
	"blog/app/verify/rpc/rpcClient"
	"blog/common/response"
	"blog/common/response/errx"
	"context"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Config    config.Config
	VerifyRpc rpcClient.Rpc
}

func NewAuthMiddleware(c config.Config, rpc rpcClient.Rpc) *AuthMiddleware {
	return &AuthMiddleware{
		Config:    c,
		VerifyRpc: rpc,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		authRes, err := m.VerifyRpc.Auth(context.Background(), &rpcClient.AuthReq{
			Server: "user",
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
		r.Header.Add("user_id", authRes.Key)
		next(w, r)
	}
}
