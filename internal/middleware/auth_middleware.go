package middleware

import (
	"fmt"
	"golang_template_source/config/constant"
	"golang_template_source/internal/repository"
	"golang_template_source/internal/usecase"
	"golang_template_source/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
)

// const (
// 	authorizationHeaderKey  = "authorization"
// 	authorizationTypeBearer = "bearer"
// 	authorizationPayloadKey = "authorization_payload"
// )

type AuthMiddleware struct {
	AuthUseCase  usecase.AuthUseCase
	db           *gorm.DB
	functionRepo repository.SysFunctionRepository
}

func NewAuthMiddleware(uc usecase.AuthUseCase, db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		db:           db,
		functionRepo: repository.NewSysFunctionRepository(db),
		AuthUseCase:  uc,
	}
}

func (m *AuthMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(constant.AuthorHeader)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.NewResponse("Token không có", nil))
			return
		}

		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.NewResponse("Format token không hợp lệ", nil))
			return
		}

		tokenString := parts[1]

		token, err := m.AuthUseCase.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.NewResponse("token không hợp lệ", nil))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, utils.NewResponse("token không hợp lệ", nil))
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims["userID"])

		ctx.Next()
	}
}


func (m *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		path := c.Request.URL.Path
		
		// Check and return original path
		originalPath, err := m.functionRepo.CheckAndReturnOriginalPath(path, constant.PATH_PREFIX)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				utils.NewResponse("Failed to process path", nil))
			c.Abort()
			return
		}

		userIDInterface, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse("User ID not found", nil))
			return
		}
		fmt.Println("userIDInterface ", userIDInterface)

		userIDStr, ok := userIDInterface.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse("Invalid User ID format", nil))
			return
		}
		userID, err := strconv.Atoi(userIDStr) // Chuyển string sang int
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse("Invalid User ID format", nil))
			return
		}

		// Truyền userID vào hàm
		isAuth, err := m.functionRepo.IsAuthentication(userID, originalPath)
		if err != nil || !isAuth {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				utils.NewResponse("ERROR_FORBIDDEN_USER_ACCESS", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
