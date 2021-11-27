package result

const (
	//请求处理成功
	OK = 0

	//数据格式无法解析
	UnknownDataFormat = 5000
	//无法找到
	NotFound = 404
	//认证错误
	ErrorAuth = 5001
	//无权限
	NoAuth      = 5002
	SystemError = 500
	SUCCESS     = 0
	ERROR       = -1
)

type Result struct {
	Code    int64
	Message string
	Data    interface{}
}

func (r *Result) isOk() bool {
	return int(r.Code) == SUCCESS
}

func (r *Result) isError() bool {
	return !r.isOk()
}
