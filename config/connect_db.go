package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitPostgreSQL initializes the PostgreSQL connection and returns the database instance
func InitPostgreSQL() *gorm.DB {
	DB = connectDB()
	return DB
}

// connectDB establishes a connection to the PostgreSQL database
func connectDB() *gorm.DB {
	dbName := os.Getenv("POSTGRESQL_NAME")
	dbPassword := os.Getenv("POSTGRESQL_PASSWORD")
	dbHost := os.Getenv("POSTGRESQL_HOST")
	dbPort := os.Getenv("POSTGRESQL_PORT")
	dbUser := os.Getenv("POSTGRESQL_USERNAME")

	// PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	sqlDB, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database handle: %v", err))
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return conn
}

// CloseConnectDB closes the database connection
func CloseConnectDB(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

// // migrateDB handles database migrations
// func migrateDB(db *gorm.DB) {
// 	type User struct {
// 		ID       uint   `gorm:"primaryKey"`
// 		Name     string `gorm:"size:100;not null"`
// 		Email    string `gorm:"unique;not null"`
// 		Password string `gorm:"not null"`
// 	}

// 	if err := db.AutoMigrate(&User{}); err != nil {
// 		panic(fmt.Sprintf("Failed to run migrations: %v", err))
// 	}
// }