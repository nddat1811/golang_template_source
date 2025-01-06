package domain

import "time"

type User struct {
	ID        int       `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}


type SysLog struct {
	ID             int `gorm:"primaryKey;autoIncrement"`
	ActionDatetime time.Time
	PathName       string
	Method         string
	IP             string
	StatusResponse int
	Response       string
	Description    string
	RequestBody    string
	RequestQuery   string
	Duration       float64
}
