package domain

import "time"

type SysFunction struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(255)"`
	Path        string    `gorm:"column:path;type:varchar(255)"`
	Regex       *string   `gorm:"column:regex;type:varchar(255)"`
	Description string    `gorm:"column:description;type:text"`
	ParentID    *int      `gorm:"column:parent_id"`
	Type        *string   `gorm:"column:type;type:varchar(50)"`
	Status      string    `gorm:"column:status;type:varchar(50)"`
	IconURL     *string   `gorm:"column:icon_url;type:varchar(255)"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy   *int      `gorm:"column:created_by"`
	UpdatedBy   *int      `gorm:"column:updated_by"`
}

func (SysFunction) TableName() string {
	return "SYS_FUNCTION"
}
