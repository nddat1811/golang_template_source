package entity

import "time"

type SysRole struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name"`
	Description *string   `gorm:"column:description"`
	Status      int       `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy   *int      `gorm:"column:created_by"`
	UpdatedBy   *int      `gorm:"column:updated_by"`
}

func (SysRole) TableName() string {
	return "SYS_ROLE"
}

type SysRoleFunction struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	RoleID     int       `gorm:"column:role_id"`
	FunctionID int       `gorm:"column:function_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy  *int      `gorm:"column:created_by"`
	UpdatedBy  *int      `gorm:"column:updated_by"`
}

func (SysRoleFunction) TableName() string {
	return "SYS_ROLE_FUNCTION"
}
