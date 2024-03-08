package main

// import (
// 	"lnovpn/internal/shell"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// )

// // 创建服务器可执行文件
// func NewServerFile() error {
// 	os.Mkdir("tmp", 0700)
// 	shell.Linux("cd template/server/;go build .;mv ./main ../../")
// 	return nil
// }

// // 创建windows可执行文件
// func NewWinFile(start, end int) error {

// 	data, err := os.ReadFile(filepath.Join("template", "client.go"))
// 	if err != nil {
// 		return err
// 	}
// 	for ; start < end; start++ {
// 		currGoFileName := cfg.SoftwareName + strconv.Itoa(start) + ".go"
// 		currGoexeName := cfg.SoftwareName + strconv.Itoa(start) + ".exe"

// 		newFileData := []byte{}
// 		copy(newFileData, data)
// 		file, err := os.OpenFile(filepath.Join("tmp", currGoFileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
// 		if err != nil {
// 			return err
// 		}

// 		file.Close()

// 		linuxShell("env GOOS=windows GOARCH=amd64 go build " + filepath.Join("tmp", currGoFileName))
// 		linuxShell("mv " + currGoexeName + " " + filepath.Join("tmp", currGoexeName))
// 	}

// 	return nil
// }
