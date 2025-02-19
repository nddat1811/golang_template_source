package main

import (
	"golang_template_source/config"
	"golang_template_source/docs"
	"golang_template_source/internal/routers"
	"golang_template_source/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization

type OK struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Ten	string
	CreatedAt time.Time `gorm:"autoCreateTime"`  // Gán giá trị mặc định khi tạo mới
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // Tự động cập nhật khi update
}

func (OK) TableName() string {
	return "OK2"
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.StartScheduler()
	conn := config.InitPostgreSQL()
	defer config.CloseConnectDB(conn)

	// conn.AutoMigrate(&OK{})

	// // // Tạo bản ghi mới
	// newRecord := OK{Name: "Test 3Name3"}
	// conn.Create(&newRecord)
	// fmt.Println("Dữ liệu ban đầu:", newRecord)

	// // Chờ 5 giây để kiểm tra cập nhật UpdatedAt
	// time.Sleep(100 * time.Second)

	// // Cập nhật Name
	// conn.Model(&newRecord).Update("Name", "Updated Name2")

	// var updatedRecord OK
	// conn.First(&updatedRecord, newRecord.ID)
	// fmt.Println("Dữ liệu sau khi cập nhật:", updatedRecord)
	//programmatically set swagger info
	docs .SwaggerInfo.Title = "Swagger Golang"
	docs.SwaggerInfo.Description = "This is a golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r :=  routers.SetupRouter(conn)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	appPort := os.Getenv("APP_PORT")
	r.Run(appPort)
}
