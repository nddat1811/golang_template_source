package domain

import "time"

type SysRole struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"type:varchar(255);default:null"`
	Description *string    `gorm:"type:text;default:null"`
	Status      *int       `gorm:"default:null"`
	CreatedAt   *time.Time `gorm:"column:created_at;default:null;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;default:null;autoUpdateTime"`
	CreatedBy   *int       `gorm:"column:created_by;default:null"`
	UpdatedBy   *int       `gorm:"column:updated_by;default:null"`
}

func (sysRole *SysRole) TableName() string {
	return "SYS_ROLE"
}

type SysRoleFunction struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	RoleID     int       `gorm:"column:role_id"`
	FunctionID int       `gorm:"column:function_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	CreatedBy  int       `gorm:"column:created_by"`
}

func (sysRoleFunction *SysRoleFunction) TableName() string {
	return "SYS_ROLE_FUNCTION"
}
