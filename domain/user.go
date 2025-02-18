package domain

import (
	"gorm.io/datatypes"
	"time"
)

type SysUser struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email     string    `gorm:"column:email;unique;not null"`
	RandomID  string    `gorm:"column:random_id;unique;not null"`
	Password  string    `gorm:"column:hash_password"`
	Phone     string    `gorm:"column:phone"`
	Name      string    `gorm:"column:full_name"`
	Status    string    `gorm:"column:status"`
	Identity  string    `gorm:"column:identity"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy *int      `gorm:"column:created_by"`
	UpdatedBy *int      `gorm:"column:updated_by"`
}

func (SysUser) TableName() string {
	return "SYS_USER"
}

type SysUserRole struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	RoleID    int       `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy *int      `gorm:"column:created_by"`
	UpdatedBy *int      `gorm:"column:updated_by"`
}

func (SysUserRole) TableName() string {
	return "SYS_USER_ROLE"
}

type UserDoc struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID        int       `gorm:"column:user_id"`
	IDCardFront   int       `gorm:"column:id_card_front"`
	IDCardBack    int       `gorm:"column:id_card_back"`
	PortraitPhoto int       `gorm:"column:portrait_photo"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy     *int      `gorm:"column:created_by"`
	UpdatedBy     *int      `gorm:"column:updated_by"`
}

func (UserDoc) TableName() string {
	return "USER_DOC"
}

type UserChangeHistory struct {
	ID           int            `gorm:"column:id;primaryKey;autoIncrement"`
	UserID       int            `gorm:"column:user_id"`
	ChangeType   string         `gorm:"column:change_type"`
	OldData      datatypes.JSON `gorm:"column:old_data"`
	NewData      datatypes.JSON `gorm:"column:new_data"`
	Status       string         `gorm:"column:status"`
	ApprovalTime *time.Time     `gorm:"column:approval_time"`
	ApprovalID   *int           `gorm:"column:approval_id"`
	Note         string         `gorm:"column:note"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy    *int           `gorm:"column:created_by"`
	UpdatedBy    *int           `gorm:"column:updated_by"`
}

func (UserChangeHistory) TableName() string {
	return "USER_CHANGE_HISTORY"
}

type UserPayment struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID        int       `gorm:"column:user_id"`
	AccountNumber string    `gorm:"column:account_number"`
	BankName      string    `gorm:"column:bank_name"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy     *int      `gorm:"column:created_by"`
	UpdatedBy     *int      `gorm:"column:updated_by"`
}

func (UserPayment) TableName() string {
	return "USER_PAYMENT"
}


type UpdateFullNameRequest struct {
	Fullname string `json:"full_name" binding:"required"`
}
// 