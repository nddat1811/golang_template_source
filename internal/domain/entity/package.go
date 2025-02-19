package entity

import "time"

type Package struct {
	ID                  int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name                string    `gorm:"column:name;type:varchar(255)"`
	Price               float64   `gorm:"column:price;not null"`
	Status              string    `gorm:"column:status;not null"`
	AffiliateCommission float64   `gorm:"column:affiliate_commission"`
	MobifoneCommission  float64   `gorm:"column:mobifone_commission"`
	AgencyCommission    float64   `gorm:"column:agency_commission"`
	Cycle               string    `gorm:"column:cycle"` //Chu kỳ
	PriorityProduct     bool      `gorm:"column:priority_product;type:boolean"`
	Benefit             string    `gorm:"column:benefit;type:text"`
	Condition           string    `gorm:"column:condition;type:text"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy           *int      `gorm:"column:created_by"`
	UpdatedBy           *int      `gorm:"column:updated_by"`
}

func (Package) TableName() string {
	return "PACKAGE"
}
