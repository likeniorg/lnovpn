package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 随机字符串
func RandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randString := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[randString.Intn(len(charset))]
	}
	return string(result)

}

// 输出提示消息，然后等待写入字符串
func Input(msg string) string {
	fmt.Println(msg)
	var tmpString string
	fmt.Scan(&tmpString)
	return tmpString
}

// 证书文件转换字符串，然后删除
func CertFileToString(path string) string {
	sstpCrt, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Remove(path)
	return string(sstpCrt)
}
