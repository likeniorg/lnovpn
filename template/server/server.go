package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"lnovpn/internal/account"
	"lnovpn/internal/shell"
	"strings"

	"github.com/gin-gonic/gin"
)

type IPcount struct {
	IP     string
	Counnt int
}

type IPsha struct {
	Sha256 string
}

func main() {

	r := gin.Default()
	// 随机生成一个复杂的URL路径作为后台链接
	r.POST(urlpath, func(ctx *gin.Context) {
		// 表单信息
		formInfo := account.Info{}
		err := ctx.ShouldBind(&formInfo)
		// 解析表单出错
		if err != nil {
			logWarnEcho(err.Error(), ctx.ClientIP())
			ctx.JSON(101, nil)
			return
		}

		// 查询是否存在该账号密码
		err = db.QueryRow("select * from accountInfo where userName = ? and password = ?", formInfo.UserName, formInfo.Password).Scan()
		if err != nil {
			if err == sql.ErrNoRows {
				strBuild := strings.Builder{}
				strBuild.WriteString("未找到用户或密码错误: ")
				strBuild.WriteString(formInfo.UserName)

				logWarnEcho(strBuild.String(), ctx.ClientIP())
				ctx.JSON(102, nil)
			} else {
				logWarnEcho(err.Error(), ctx.ClientIP())
			}
		}

		ipCount := IPcount{}
		err = db.QueryRow("SELECT count FROM ipCount where ip = ?", ctx.ClientIP()).Scan(ipCount.Counnt)
		if err != nil {
			// 如果不存在该IP，放行防火墙
			if err == sql.ErrNoRows {
				db.Exec("INSERT INTO ipCount VALUES(?,?)", ctx.ClientIP(), 1)
				shell.AllowIP(ctx.ClientIP())
			} else {
				// 查找IP失败
				logWarnEcho(err.Error(), ctx.ClientIP())
				ctx.JSON(103, nil)
			}
		}
		// IP防火墙已被放行过，更新计数器
		db.Exec("UPDATE ipCount SET count = ? WHERE ip = ?", ipCount.Counnt+1, ctx.ClientIP())
		ctx.JSON(200, nil)
	})

	r.POST(keepURL, func(ctx *gin.Context) {
		ipsha := IPsha{}
		sha := sha256.New()
		ipCount := IPcount{}

		err := ctx.ShouldBind(&ipsha)
		if err != nil {
			logWarnEcho("读取客户端断开连接HASH失败", ctx.ClientIP())
			ctx.JSON(104, nil)
		}

		if ctx.ClientIP() == hex.EncodeToString(sha.Sum([]byte(ipsha.Sha256))) {
			err := db.QueryRow("SELECT count FROM ipCount WHERE = ?", ctx.ClientIP()).Scan(ipCount.IP)
			if err != nil {
				if err == sql.ErrNoRows {
					logWarnEcho("数据库中不存在此IP，无法被销毁", ctx.ClientIP())
					ctx.JSON(105, nil)
				} else {
					logWarnEcho("销毁IP时查询IP操作失败"+err.Error(), ctx.ClientIP())
					ctx.JSON(106, nil)
				}
			}
		}

		shell.RemoveIP(ctx.ClientIP())
		ctx.JSON(201, nil)
	})
}

var urlpath = "/almxaomxnauada"
var keepURL = "/xkoaxdas"
