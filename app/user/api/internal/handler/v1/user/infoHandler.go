package user

import (
	"blog/app/user/api/internal/logic/v1/user"
	"blog/app/user/api/internal/svc"
	"blog/app/user/api/internal/types"
	"blog/common/response"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"reflect"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		//解析参数
		httpx.Parse(r, &req)
		//验证器
		uni := ut.New(zh.New())
		trans, _ := uni.GetTranslator("zh")
		validate := validator.New()
		zhTrans.RegisterDefaultTranslations(validate, trans)
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("label")
		})
		errs := validate.Struct(req)
		if errs != nil {
			first := errs.(validator.ValidationErrors)[0]
			response.ParamError(w, first.Translate(trans))
			return
		}
		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)

		}
	}
}
