package usecase

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"golang_template_source/internal/domain"
	"golang_template_source/internal/domain/dto"
	"golang_template_source/internal/repository"
	"html/template"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type AuthUseCase interface {
	Login(email, password string) (dto.TokenResponse, error)
	Register(user *domain.SysUser) error
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(refreshTokenString string) (dto.TokenResponse, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
	functionRepo repository.SysFunctionRepository
}

var jwtKey = os.Getenv("API_KEY")

func NewAuthUseCase(userRepo repository.UserRepository, functionRepo repository.SysFunctionRepository) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
		functionRepo: functionRepo,
	}
}

func (a *authUseCase) issueTokens(userID string) (dto.TokenResponse, error) {
	// Phát hành Access Token (hết hạn trong 1 ngày)
	accessTokenExpiration := time.Now().Add(24 * time.Hour)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    accessTokenExpiration.Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		return dto.TokenResponse{}, err
	}

	// Phát hành Refresh Token (hết hạn trong 7 ngày)
	refreshTokenExpiration := time.Now().Add(7 * 24 * time.Hour)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    refreshTokenExpiration.Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(jwtKey))
	if err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken:      accessTokenString,
		RefreshToken:     refreshTokenString,
	}, nil
}

func (a *authUseCase) Login(email, password string) (dto.TokenResponse, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return dto.TokenResponse{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return dto.TokenResponse{}, errors.New("invalid credentials222")
	}

	return a.issueTokens(strconv.Itoa(user.ID))
}

func SendActivationEmail(smtpServer string, smtpPort int, emailSender, emailPassword, receiversEmail, token string) error {
	dialer := gomail.NewDialer(smtpServer, smtpPort, emailSender, emailPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Bỏ qua xác thực TLS nếu cần

	activationLink := fmt.Sprintf("http://localhost:8080/activate?token=%s", token)

	tmpl, err := template.New("activationEmail").Parse(`
		<html>
			<body>
				<p>Xin chào,</p>
				<p>Vui lòng nhấp vào liên kết dưới đây để kích hoạt tài khoản của bạn:</p>
				<a href="{{.Link}}">Kích hoạt tài khoản</a>
				<p>Liên kết này sẽ hết hạn sau 1 giờ.</p>
			</body>
		</html>`)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, struct{ Link string }{Link: activationLink})
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", emailSender)
	msg.SetHeader("To", receiversEmail)
	msg.SetHeader("Subject", "Kích hoạt tài khoản của bạn")
	msg.SetBody("text/html", body.String())

	// Gửi email
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}


func (a *authUseCase) Register(user *domain.SysUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Status = "0"
	// createdUser, errCreate := a.userRepo.Create(user)
	// if errCreate!= nil {
    //     return errCreate
    // }

	// Tạo token kích hoạt
	token, err := a.issueTokens(strconv.Itoa(3))
	if err != nil {
		return err
	}
	smtpServer := ""
	smtpPort := 25
	emailSender := "product.mbf2@mobifone.vn"
	emailPassword := ""
	receiversEmail := "yongpalkim1811@gmail.com"

	// Gửi email kích hoạt
	err = SendActivationEmail(smtpServer, smtpPort, emailSender, emailPassword, receiversEmail, token.AccessToken)
	if err != nil {
		return err
	}
	return nil
}

func (a *authUseCase) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
}


func (a *authUseCase) RefreshToken(refreshTokenString string) (dto.TokenResponse, error) {
	// Parse và xác thực Refresh Token
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return dto.TokenResponse{}, errors.New("invalid refresh token")
	}

	// Lấy thông tin userID từ claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return dto.TokenResponse{}, errors.New("invalid refresh token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return dto.TokenResponse{}, errors.New("invalid refresh token payload")
	}

	// Phát hành cặp token mới cho user
	return a.issueTokens(userID)
}