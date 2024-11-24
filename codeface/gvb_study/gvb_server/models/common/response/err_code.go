package response

type ErrorCode int

const (
	SettingsError ErrorCode = 10001
	ArgumentError ErrorCode = 10002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
