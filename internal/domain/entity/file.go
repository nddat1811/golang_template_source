package entity

import (
	"time"
)

type SysFile struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	ShareLink string    `gorm:"column:share_link"`
	Type      string    `gorm:"column:type"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	CreatedBy *int      `gorm:"column:created_by"`
}

func (SysFile) TableName() string {
	return "SYS_FILE"
}
