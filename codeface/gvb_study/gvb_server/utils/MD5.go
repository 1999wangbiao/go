package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func GenerateMD5FromFileHeader(fileHeader *multipart.FileHeader) (string, error) {
	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建 MD5 哈希对象
	hash := md5.New()

	// 读取文件内容并写入到 MD5 哈希对象
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// 获取 MD5 字节值
	md5Hash := hash.Sum(nil)

	// 返回 MD5 的十六进制表示
	return fmt.Sprintf("%x", md5Hash), nil
}
