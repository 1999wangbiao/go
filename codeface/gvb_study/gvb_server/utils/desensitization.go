package utils

import "strings"

// DesensitizationEmail 邮件加密
func DesensitizationEmail(email string) string {
	split := strings.Split(email, "@")
	if len(split) != 2 {
		return ""
	}
	return split[0][:1] + "*****@" + split[1]
}

// DesensitizationTel 邮件加密
func DesensitizationTel(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "*****" + tel[7:]
}
