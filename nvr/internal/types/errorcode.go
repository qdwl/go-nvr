package types

type ErrorCode int64

const (
	RESTFUL_ERR_OK                        ErrorCode = 0
	RESTFUL_ERR_AUTH_EXPIRE               ErrorCode = 2000
	RESTFUL_ERR_AUTH_FAILED               ErrorCode = 2001
	RESTFUL_ERR_UNKNOW                    ErrorCode = 9000 //未知错误
	RESTFUL_ERR_NOLOGIN                   ErrorCode = 9001 //用户未登陆
	RESTFUL_ERR_PERMISSION                ErrorCode = 9002 //无权限
	RESTFUL_ERR_HTTP_HEAD_MISS_PARAM      ErrorCode = 9003 //HTTP请求头缺少参数
	RESTFUL_ERR_PARSE_BODY_FAIL           ErrorCode = 9004 //请求http包体解析失败
	RESTFUL_ERR_PARAMETER_FAILED          ErrorCode = 9005 //参数错误
	RESTFUL_ERR_DATABASE_OPERATION_FAILED ErrorCode = 9006 //操作数据库失败
	RESTFUL_ERR_STARTUP_GB28181_SERVER    ErrorCode = 9007 //重启国标服务失败
)

func (ec ErrorCode) String() string {
	switch ec {
	case RESTFUL_ERR_OK:
		return "success"
	case RESTFUL_ERR_AUTH_EXPIRE:
		return "authorization expiration"
	case RESTFUL_ERR_AUTH_FAILED:
		return "authorization failed"
	case RESTFUL_ERR_UNKNOW:
		return "unknow error"
	case RESTFUL_ERR_NOLOGIN:
		return "no login"
	case RESTFUL_ERR_PERMISSION:
		return "no permissions"
	case RESTFUL_ERR_HTTP_HEAD_MISS_PARAM:
		return "parameter error"
	case RESTFUL_ERR_PARSE_BODY_FAIL:
		return "parse body failed"
	case RESTFUL_ERR_DATABASE_OPERATION_FAILED:
		return "database operation failed"
	case RESTFUL_ERR_STARTUP_GB28181_SERVER:
		return "startup gb28181 server failed"
	}

	return ""
}
