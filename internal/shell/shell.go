package shell

import (
	"os/exec"
	"strings"
)

// 放行防火墙IP
func AllowIP(ip string) {
	strBuild := strings.Builder{}
	strBuild.WriteString(`firewall-cmd  --zone=public --add-rich-rule="rule family="ipv4" source address="`)
	strBuild.WriteString(ip)
	strBuild.WriteString(`" accept"`)

	exec.Command("bash", "-c", strBuild.String())
}

// 停止放行防火墙IP
func RemoveIP(ip string) {
	strBuild := strings.Builder{}
	strBuild.WriteString(`firewall-cmd  --zone=public --remove-rich-rule="rule family="ipv4" source address="`)
	strBuild.WriteString(ip)
	strBuild.WriteString(`" accept"`)

	exec.Command("bash", "-c", strBuild.String())
}

// Linux Shell
func Linux(dst string) error {
	cmd := exec.Command("bash", "-c", dst)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}
