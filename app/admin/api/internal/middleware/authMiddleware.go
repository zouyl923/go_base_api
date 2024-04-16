package middleware

import (
	"blog/common/helper"
	"blog/common/response"
	"blog/common/response/errx"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		claims, err := helper.ParseToken(token)
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
		r.Header.Set("AdminId", claims.Key)
		next(w, r)
	}
}
