package domain

import "time"

type SysLog struct {
	ID             int `gorm:"primaryKey;autoIncrement"`
	ActionDatetime time.Time
	PathName       string
	Method         string
	IP             string
	StatusResponse int
	Response       string
	Description    string
	RequestBody    string
	RequestQuery   string
	Duration       float64
}

func (SysLog) TableName() string {
	return "SYS_LOG"
}
