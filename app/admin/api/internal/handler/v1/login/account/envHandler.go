package account

import (
	"blog/app/admin/api/internal/logic/v1/login/account"
	"blog/app/admin/api/internal/svc"
	"blog/common/response"
	"net/http"
)

func EnvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewEnvLogic(r.Context(), svcCtx)
		resp, err := l.Env(r)
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
