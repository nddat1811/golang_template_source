package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"golang_template_source/config"
	"golang_template_source/docs"
	"golang_template_source/routers"
	"golang_template_source/utils"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)


// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.StartScheduler()
	conn := config.InitPostgreSQL()
	defer config.CloseConnectDB(conn)

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
