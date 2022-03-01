package exception

import (
	"net/http"
	"time"
)

// Exception 对客户端返回错误类型的标准定义
type Exception struct {
	Timestamp time.Time `json:"timestamp"` // 时间戳
	Code      int       `json:"code"`      // 错误码
	Status    int       `json:"status"`    // http状态
	Message   string    `json:"message"`   // 错误消息
	Path      string    `json:"path"`      // 调用路由
}

func (e *Exception) Error() string {
	return e.Message
}

func newException(code int, status int, message string) *Exception {
	return &Exception{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

const (
	// CodeServerError 系统错误
	CodeServerError = -1
	// CodeParameterError 参数错误
	CodeParameterError = -2
	// CodeAuthError 权限错误
	CodeAuthError = -3
	// CodeUnkownError 未知错误
	CodeUnkownError = -4
)

// ServerError returns a standard 500 error
func ServerError() *Exception {
	return newException(
		CodeServerError,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

// ParameterError returns a parameter error
func ParameterError() *Exception {
	return newException(
		CodeParameterError,
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
	)
}

// ParameterError returns a parameter error
func AuthError() *Exception {
	return newException(
		CodeAuthError,
		http.StatusUnauthorized,
		http.StatusText(http.StatusUnauthorized),
	)
}

// ParameterError returns a parameter error
func UnknownError(message string) *Exception {
	return newException(
		CodeUnkownError,
		http.StatusForbidden,
		message,
	)
}
