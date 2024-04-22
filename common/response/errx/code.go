package errx

const (
	Success    int = 0  //成功
	Error      int = -1 //通用失败
	LoginError int = -2 //登录错误
	ParamError int = -3 //参数错误

	NotFundError int = -4 //查询失败
	DeleteError  int = -5 //删除失败
	UpdateError  int = -6 //更新失败
	MessageError int = -8 //自定义业务错误
	UnKnowError  int = -9 //未知错误
	//admin模块 101开头 01 是具体业务
	AdminNotFound int = 10101 //账户或者密码错误

)
