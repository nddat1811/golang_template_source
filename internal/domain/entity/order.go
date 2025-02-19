package entity

import (
	"time"
)

type Order struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID        int       `gorm:"column:user_id;not null"`
	CustomerPhone string    `gorm:"column:customer_phone"`
	Status        string    `gorm:"column:status"`
	TotalPrice    float64   `gorm:"column:total_price"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy     *int      `gorm:"column:created_by"`
	UpdatedBy     *int      `gorm:"column:updated_by"`
}

func (Order) TableName() string {
	return "ORDER"
}

// Struct GORM cho bảng ORDER_PACKAGE
type OrderPackage struct {
	OrderID             int       `gorm:"column:order_id;primaryKey"`
	PackageID           int       `gorm:"column:package_id;primaryKey"`
	Price               float64   `gorm:"column:price"`
	AffiliateCommission float64   `gorm:"column:affiliate_commission"`
	MobifoneCommission  float64   `gorm:"column:mobifone_commission"`
	AgencyCommission    float64   `gorm:"column:agency_commission"`
	Status              string    `gorm:"column:status"` // đã rút tiền, bình thường
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy           *int      `gorm:"column:created_by"` // Có thể NULL
	UpdatedBy           *int      `gorm:"column:updated_by"` // Có thể NULL
}

func (OrderPackage) TableName() string {
	return "ORDER_PACKAGE"
}
