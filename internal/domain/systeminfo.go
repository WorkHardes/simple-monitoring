package domain

type SystemInfo struct {
	SysDescr  string
	SysUpTime WorkTime
	SysName   string
}

func NewSystemInfo() SystemInfo {
	return SystemInfo{}
}
