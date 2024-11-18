package res

type ErrorCode int

const (
	SettringsError ErrorCode = 10001
)

var (
	ErrorMap = map[ErrorCode]string{
		SettringsError: "系统错误",
	}
)
