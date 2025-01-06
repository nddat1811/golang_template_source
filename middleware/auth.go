package middleware

import (
	"golang_template_source/config/constant"
	"golang_template_source/usecase"
	"golang_template_source/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// const (
// 	authorizationHeaderKey  = "authorization"
// 	authorizationTypeBearer = "bearer"
// 	authorizationPayloadKey = "authorization_payload"
// )

type AuthMiddleware struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthMiddleware(uc usecase.AuthUseCase) *AuthMiddleware{
	return &AuthMiddleware{AuthUseCase: uc}
}


func (m *AuthMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth := service.NewAuthService()

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
			ctx.JSON(http.StatusUnauthorized, utils.NewResponse("Bad request", nil))
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims["userID"])
		ctx.Next()
	}
}