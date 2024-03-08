package account

import "time"

// 账号信息
type Info struct {
	UserName      string `form:"userName" binding:"required"`
	Password      string `form:"password" binding:"required"`
	HostID        string `form:"hostId" binding:"required"`
	VuserName     string
	Vpassword     string
	Note          string
	StartTime     time.Time
	EndTime       time.Time
	LastLoginTime time.Time
}
