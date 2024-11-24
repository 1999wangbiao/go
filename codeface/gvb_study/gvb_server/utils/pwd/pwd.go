package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd hash 密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePassword 验证密码  hash之后的密码 输入的密码
func ComparePassword(hashedPwd string, plainPwd string) bool {
	bytesHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(bytesHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
