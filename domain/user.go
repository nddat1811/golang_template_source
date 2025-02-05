package domain

import "time"

type SysUser struct {
	ID        int       `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:full_name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	Password  string    `gorm:"column:hash_password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}


func (sysUser *SysUser) TableName() string {
	return "SYS_USER"
}

type SysUserRole struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	RoleID    int       `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	CreatedBy int       `gorm:"column:created_by"`
}

func (sysUserRole *SysUserRole) TableName() string {
	return "SYS_USER_ROLE"
}


type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateFullNameRequest  struct {
	Fullname    string `json:"full_name" binding:"required"`
}

type LoginRequest  struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
