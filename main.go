package main

import (
	"fmt"
	"lnovpn/internal/util"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}
func main() {
	// 检测是否存在配置文件
	_, err := os.Stat(".env")
	if err != nil {
		if is := util.Input("未检测配置文件，是否导入(y/n)"); is == "y" {
			fmt.Println("请将.env文件存放到项目根目录后运行")
			os.Exit(0)
		} else {
			if is == "n" {
				makeConf()
			} else {
				fmt.Println("输入不符合要求，退出程序")
				os.Exit(1)
			}
		}
	}

}
