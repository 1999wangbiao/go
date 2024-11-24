package response

import "gvb_server/models/system"

type SysUserResponse struct {
	User system.UserModel `json:"user"`
}

type LoginResponse struct {
	User      system.UserModel `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}
