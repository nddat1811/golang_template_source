package routers

import (
	"os"
	"time"

	"golang_template_source/controller"
	"golang_template_source/middleware"
	"golang_template_source/repository"
	"golang_template_source/usecase"
	"golang_template_source/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func SetupRouter(conn *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_ALLOWED_ORIGINS")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middleware.LogMiddleware())

	userRepo := repository.NewUserRepository(conn)
	functionRepo := repository.NewSysFunctionRepository(conn)

	authUseCase := usecase.NewAuthUseCase(userRepo, functionRepo)

	authController := controller.NewAuthController(authUseCase)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/refresh", authController.RefreshToken)
	}

	userUseCase := usecase.NewUserUseCase(userRepo)

	user := controller.NewUserController(userUseCase)
	
	authMiddleware := middleware.NewAuthMiddleware(authUseCase, conn)

	protected := router.Group("/")
	protected.Use(authMiddleware.TokenAuthMiddleware())
	protected.Use(authMiddleware.Middleware())
	{
		protected.GET("/users/:id", user.GetUserByID)
		protected.GET("/users", user.GetAllUsers)
	}
	router.GET("/users/export", user.ExportUsersToExcel)
	router.GET("/users/export-template", user.ExportUsersToTemplate)

	manager := utils.NewConnectionManager()
	socketController := controller.NewWebSocketController(manager)
	router.GET("/socket/ws/:room", socketController.HandleWebSocket)

	return router
}
