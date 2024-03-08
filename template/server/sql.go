package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	DB, err := sql.Open("mysql", "root:k@tcp(127.0.0.1:3306)/dbname?charset=utf8")
	if err != nil {
		logErrEcho(err.Error(), "")
	}
	db = DB
	if err = db.Ping(); err != nil {
		logErrEcho(err.Error(), "数据库连接失败")
	}
	logInfoEcho("数据库连接成功", "")
	db.Exec(`CREATE DATABASE IF NOT EXISTS lnovpn`)
	db.Exec(`CREATE TABLE IF NOT EXISTS accountInfo(
		userName char(11) PRIMARY KEY,
		password char(11) not null,
		hostID varchar(255) not null,
		vusername char(11) not null,
		vpassword char(11) not null,
		note varchar(255) default null,
		startTime varchar(255) not null,
		endTime varchar(255) not null,
		lastLoginTime varchar(255) not null
	)`)
	db.Exec(`CREATE TABLE  IF NOT EXISTS ipCount(
		ip varchar(255) PRIMARY KEY,
		count int default 0
	)`)

}

var db *sql.DB
