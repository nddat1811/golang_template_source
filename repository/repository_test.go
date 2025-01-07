package repository_test

import (

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "database/sql/driver"
	"database/sql"
	"time"
)

var (
	createdAt = time.Now()
	updatedAt = time.Now()
)

// type AnyTime struct{}

// func (a AnyTime) Match(v driver.Value) bool {
// 	_, ok := v.(time.Time)

// 	return ok
// }

// Setup mock database for testing
func NewConnManager() (*sql.DB, *gorm.DB, sqlmock.Sqlmock, error) {
	// Tạo mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	// Cấu hình PostgreSQL driver cho GORM
	gdb, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn:                 db,                   // Sử dụng mock database
			PreferSimpleProtocol: true,                // Đơn giản hóa giao thức giao tiếp
		}),
		&gorm.Config{},
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return db, gdb, mock, nil
}