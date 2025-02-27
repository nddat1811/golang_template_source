package controller

import (
	"golang_template_source/internal/domain"
	"golang_template_source/internal/domain/dto"
	"golang_template_source/internal/usecase"
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
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} string
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Bad request", nil))
		return
	}

	token, err := c.authUseCase.Login(request.Email, request.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK,utils.NewResponse("ok", token))
}


func (c *AuthController) Login2(ctx *gin.Context) {
	var request dto.LoginRequest

	// Bước 1: Bind dữ liệu từ request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Bad request", nil))
		return
	}

	// Bước 2: Lấy token reCAPTCHA từ request
	recaptchaToken := request.ReCaptcha
	if recaptchaToken == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse("Missing reCAPTCHA token", nil))
		return
	}

	// Bước 3: Xác minh reCAPTCHA v3
	success, score, err := utils.VerifyRecaptcha(recaptchaToken)
	if err != nil || !success {
		ctx.JSON(http.StatusUnauthorized, utils.NewResponse("reCAPTCHA verification failed", nil))
		return
	}

	// Bước 4: Kiểm tra điểm số reCAPTCHA (0.5 là ngưỡng tối thiểu)
	if score < 0.5 {
		ctx.JSON(http.StatusForbidden, utils.NewResponse("Low reCAPTCHA score, possible bot", nil))
		return
	}

	// Bước 5: Thực hiện đăng nhập nếu xác thực reCAPTCHA thành công
	token, err := c.authUseCase.Login(request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse("ok", token))
}

// User Register  godoc
// @Summary User Register
// @Description User Register
// @Accept  json
// @Produce  json
// @Param body body dto.UserRegisterInput false "body"
// @Success 200 {object} string
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var user dto.UserRegisterInput

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

// Get Refresh Token  godoc
// @Summary Get Refresh Token
// @Description Get Refresh Token
// @Accept  json
// @Produce  json
// @Param body body RefreshToken false "body"
// @Success 200 {object} string
// @Router /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request dto.RefreshTokenRequest

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