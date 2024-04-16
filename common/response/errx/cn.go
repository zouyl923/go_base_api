package errx

var cn = map[int]string{
	Success:       "操作成功！",
	Error:         "操作失败！",
	LoginError:    "登录失效",
	ParamError:    "参数错误",
	AdminNotFound: "账号或者密码错误！",
}

func GetCnMessage(code int) string {
	msg, ok := cn[code]
	if !ok {
		msg = cn[Error]
	}
	return msg
}
