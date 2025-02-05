package domain

import "time"

type SysFunction struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255)"`
	Path        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:text"`
	ParentID    *int      `gorm:"column:parent_id"`
	Type        *string   `gorm:"type:varchar(50)"`
	Status      string    `gorm:"type:varchar(50)"`
	IconURL     *string   `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy   int       `gorm:"column:created_by"`
	UpdatedBy   *int      `gorm:"column:updated_by"`
	Regex       *string   `gorm:"type:varchar(255)"`
}

func (sysFunction *SysFunction) TableName() string {
	return "SYS_FUNCTION"
}