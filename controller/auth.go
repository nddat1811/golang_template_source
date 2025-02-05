package controller

import (
	"fmt"
	"golang_template_source/domain"
	"golang_template_source/usecase"
	"golang_template_source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}


// User Login  godoc
// @Summary User Login
// @Description User Login
// @Accept  json
// @Produce  json
// @Param request body domain.LoginRequest true "Login Request"
// @Success 200 {object} string
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request domain.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Bad request", nil))
		return
	}

	token, err := c.authUseCase.Login(request.Email, request.Password)
	fmt.Println("s:", err)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewResponse("Bad request333", nil))
		return
	}

	ctx.JSON(http.StatusOK,utils.NewResponse("ok", token))
}

type UserRegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// User Register  godoc
// @Summary User Register
// @Description User Register
// @Accept  json
// @Produce  json
// @Param body body UserRegisterInput false "body"
// @Success 200 {object} string
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var user UserRegisterInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Bad request", nil))
		return
	}

	err := c.authUseCase.Register(&domain.SysUser{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse("Bad request", nil))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewResponse("ok", nil))
}


type RefreshToken struct {
	RefreshToken    string `json:"refresh_token" binding:"required"`
}

// Get Refresh Token  godoc
// @Summary Get Refresh Token
// @Description Get Refresh Token
// @Accept  json
// @Produce  json
// @Param body body RefreshToken false "body"
// @Success 200 {object} string
// @Router /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request RefreshToken

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Bad request", nil))
		return
	}

	token, err := c.authUseCase.RefreshToken(request.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.NewResponse("Bad request", nil))
		return
	}

	ctx.JSON(http.StatusOK,utils.NewResponse("ok", token))
}