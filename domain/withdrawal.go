package domain

import (
	"time"
)

type Withdrawal struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int       `gorm:"column:user_id"` // Người yêu cầu rút tiền
	RequestCode string    `gorm:"column:request_code;unique;not null"`
	Status      string    `gorm:"column:status"`
	Amount      float64   `gorm:"column:amount"`
	ProcessedBy *int      `gorm:"column:processed_by"` // Người xử lý (Admin) - có thể NULL
	Note        string    `gorm:"column:note;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy   *int      `gorm:"column:created_by"`
	UpdatedBy   *int      `gorm:"column:updated_by"`
}

func (Withdrawal) TableName() string {
	return "WITHDRAWAL"
}

type WithdrawalOrder struct {
	WithdrawalID     int     `gorm:"column:withdrawal_id;primaryKey"`
	OrderID          int     `gorm:"column:order_id;primaryKey"`
	CommissionAmount float64 `gorm:"column:commission_amount"`
}

func (WithdrawalOrder) TableName() string {
	return "WITHDRAWAL_ORDER"
}
