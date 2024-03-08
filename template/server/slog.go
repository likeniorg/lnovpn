package main

import (
	"log/slog"
	"os"
)

var slogs = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// 记录正常访问的日志
func logInfoEcho(ip, info string) {
	slogs.Info(info, "ip", ip)

}

// 记录警告日志
func logWarnEcho(warn, ip string) {
	slogs.Warn(warn, "ip", ip)

}

// 记录严重错误
func logErrEcho(err, ip string) {
	slogs.Error(err, "ip", ip)
	os.Exit(1)
}
