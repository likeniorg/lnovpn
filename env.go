package main

import (
	"encoding/json"
	"lnovpn/internal/shell"
	"lnovpn/internal/util"
	"os"
)

// 配置结构体
type Conf struct {
	Domain        string
	DomainAndPort string
	SSTPCrt       string
	SSTPKey       string
	ServerCrt     string
	ServerKey     string
	ClientCrt     string
	ClientKey     string
	ServerCsr     string
	ClientCsr     string
	SSTPCrtPath   string
	SSTPKeyPath   string
	ServerCrtPath string
	ServerKeyPath string
	ClientCrtPath string
	ClientKeyPath string
	ServerCsrPath string
	ClientCsrPath string
	Key           string
	LoginURL      string
	KeepURL       string
	DBusername    string
	DBPassword    string
	DBname        string
}

// 制造配置文件
func makeConf() {
	conf := Conf{
		// 主要修改你的域名及服务器对应监听端口
		Domain:        "",
		DomainAndPort: "",
		// 以下配置建议非必要不改动
		SSTPCrt:       "",
		SSTPKey:       "",
		ServerCsr:     "",
		ClientCsr:     "",
		ServerCrt:     "",
		ServerKey:     "",
		ClientCrt:     "",
		ClientKey:     "",
		SSTPCrtPath:   "conf/cert/sstp.crt",
		SSTPKeyPath:   "conf/cert/sstp.key",
		ServerCsrPath: "conf/cert/server.csr",
		ClientCsrPath: "conf/cert/server.csr",
		ServerCrtPath: "conf/cert/server.crt",
		ServerKeyPath: "conf/cert/server.key",
		ClientCrtPath: "conf/cert/client.crt",
		ClientKeyPath: "conf/cert/client.key",
		Key:           util.RandString(30),
		LoginURL:      "/" + util.RandString(15),
		KeepURL:       "/" + util.RandString(15),
		DBname:        "lnovpn",
	}

	// 终端输入特定信息
	conf.Domain = util.Input("输入你的服务器域名 例如:\nlno.likeni.org")
	conf.DomainAndPort = conf.Domain + ":" + util.Input("指定程序端口 例如:\n59812")
	conf.DBusername = util.Input("输入Mysql数据库用户名 例如\nroot")
	conf.DBPassword = util.Input("输入Mysql数据库用户名 例如\npassword")

	// 生成SSTP证书,将其写入到.env文件后删除原本文件
	shell.Linux(`openssl req -x509 -nodes -days 365 -newkey rsa:3072 -keyout conf/cert/sstp.key -out conf/cert/sstp.crt -subj "/C=US/ST=California/L=SanFrancisco/O=likeniorg/OU=a/CN=` + conf.Domain + `"`)
	conf.SSTPCrt = util.CertFileToString("conf/cert/sstp.crt")
	conf.SSTPKey = util.CertFileToString("conf/cert/sstp.key")

	// 生成服务器https证书
	shell.Linux(`openssl req -newkey rsa:3072 -nodes -keyout conf/cert/server.key -out conf/cert/server.csr -subj "/C=US/ST=New York/L=New York/O=lno/OU=lno/CN=` + conf.Domain + `"`)
	// shell.Linux(`openssl req -newkey rsa:3072 -nodes -keyout conf/cert/server.key -out conf/cert/server.csr`)
	shell.Linux(`openssl req -newkey rsa:3072 -nodes -keyout conf/cert/client.key -out conf/cert/client.csr -subj "/C=US/ST=New York/L=New York/O=lno/OU=lno/CN=` + conf.Domain + `"`)
	shell.Linux(`openssl req -newkey rsa:3072 -nodes -keyout conf/cert/client.key -out conf/cert/client.csr`)
	shell.Linux(`openssl x509 -req -in conf/cert/server.csr -signkey conf/cert/server.key -out conf/cert/server.crt`)
	shell.Linux(`openssl x509 -req -in conf/cert/client.csr -signkey conf/cert/client.key -out conf/cert/client.crt`)

	// 将文件转换为字符串保存后删除原本文件路径
	conf.ServerCrt = util.CertFileToString("conf/cert/server.crt")
	conf.ServerKey = util.CertFileToString("conf/cert/server.key")
	conf.ClientCrt = util.CertFileToString("conf/cert/client.crt")
	conf.ClientKey = util.CertFileToString("conf/cert/client.key")
	conf.ServerCsr = util.CertFileToString("conf/cert/server.csr")
	conf.ClientCsr = util.CertFileToString("conf/cert/client.csr")

	// 转换为JSON数据后写入文件
	json, _ := json.MarshalIndent(conf, " ", "	")
	os.WriteFile(".env", json, 0400)

}
