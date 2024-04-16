package errx

const (
	Success    int = 0  //成功
	Error      int = -1 //通用失败
	LoginError int = -2 //参数错误
	ParamError int = -3 //参数错误
	//admin模块 101开头 01 是具体业务
	AdminNotFound int = 10101 //账户或者密码错误
)
