package ctype

import "github.com/goccy/go-json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGit   SignStatus = 2
	SignEmail SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ用户"

	case SignGit:
		return "Git用户"

	case SignEmail:
		return "Email"

	default:
		return "其他用户"
	}
}
