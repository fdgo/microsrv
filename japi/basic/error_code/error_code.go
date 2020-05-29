package error_code

const (
	//-----------------------------------------------------------------------------
	//调用成功
	SUCCESS             = 200
	SUCCESS_MSG         = "成功! "
	SUCCESS_NOVALUE     = SUCCESS + 1
	SUCCESS_NOVALUE_MSG = "成功, 数据为空 !"

	ERROR      = 400
	ERROR_MSG  = "失败! "
	ERROR_DATA = "fail!"

	//-----------------------------------------------------------------------------
	ERROR_COMMON_BASE = 20000
	//参数错误
	ERROR_COMMON_PARAM     = ERROR_COMMON_BASE + 1
	ERROR_COMMON_PARAM_MSG = "请求参数有误: "
	//正则表达式错误
	ERROR_COMMON_REGEXP_USERACCOUNT     = ERROR_COMMON_BASE + 2
	ERROR_COMMON_REGEXP_USERACCOUNT_MSG = "账号格式有误( 6-16位数字和字母,可用特殊字符@#$& ): "

	ERROR_COMMON_REGEXP_PASSWORD     = ERROR_COMMON_BASE + 3
	ERROR_COMMON_REGEXP_PASSWORD_MSG = "密码格式有误( 6-16位数字和字母,可用特殊字符@#$& ): "

	ERROR_COMMON_REGEXP_MOBILE     = ERROR_COMMON_BASE + 4
	ERROR_COMMON_REGEXP_MOBILE_MSG = "手机号格式有误: "

	ERROR_COMMON_REGEXP_EMAIL     = ERROR_COMMON_BASE + 5
	ERROR_COMMON_REGEXP_EMAIL_MSG = "邮件地址格式有误: "

	//redis 链接失败：
	ERROR_COMMON_REDISCON     = ERROR_COMMON_BASE + 6
	ERROR_COMMON_REDISCON_MSG = "redis 连接失败: "

	//mysql 链接失败：
	ERROR_COMMON_MYSQLCON     = ERROR_COMMON_BASE + 7
	ERROR_COMMON_MYSQLCON_MSG = "mysql 连接失败: "

	//nsq   链接失败：
	ERROR_COMMON_NSQCON     = ERROR_COMMON_BASE + 8
	ERROR_COMMON_NSQCON_MSG = "nsq 连接失败: "

	//没有权限
	ERROR_COMMON_AUTH     = ERROR_COMMON_BASE + 9
	ERROR_COMMON_AUTH_MSG = "没有权限: "

	ERROR_COMMON_CODE_EXPIRED     = ERROR_COMMON_BASE + 10
	ERROR_COMMON_CODE_EXPIRED_MSG = "验证码过期！"

	ERROR_COMMON_CODE_WRONG     = ERROR_COMMON_BASE + 11
	ERROR_COMMON_CODE_WRONG_MSG = "验证码错误！"
	//-----------------------------------------------------------------------------

	//短信服务
	ERROR_BASESRV_BASE          = 10000                  //基础服务
	ERROR_BASESRV_CODE_WRONG    = ERROR_BASESRV_BASE + 1 //验证码错误
	ERROR_BASESRV_CODE_EXPIRED  = ERROR_BASESRV_BASE + 2 //验证码过期
	ERROR_BASESRV_CODE_NOTFOUNT = ERROR_BASESRV_BASE + 3 //验证码不存在

	//用户服务
	ERROR_USERSRV_BASE                  = 20000                  //用户服务
	ERROR_USERSRV_USER_REGISTED         = ERROR_USERSRV_BASE + 1 //用户已经注册
	ERROR_USERSRV_USER_USERNOTFOUNT     = ERROR_USERSRV_BASE + 2 //用户没找到
	ERROR_USERSRV_DEVICE_DEVICENOTFOUNT = ERROR_USERSRV_BASE + 3 //设备没有找到
	ERROR_USERSRV_DEVICE_DEVICEREGISTED = ERROR_USERSRV_BASE + 4 //设备已经注册过

	//后台服务
	ERROR_ADMINSRV_BASE = 30000


)
