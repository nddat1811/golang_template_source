package usecase

import (
	"errors"
	"golang_template_source/domain"
	"golang_template_source/repository"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(email, password string) (domain.Token, error)
	Register(user *domain.SysUser) error
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(refreshTokenString string) (domain.Token, error)
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

func (a *authUseCase) issueTokens(userID string) (domain.Token, error) {
	// Phát hành Access Token (hết hạn trong 1 ngày)
	accessTokenExpiration := time.Now().Add(24 * time.Hour)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    accessTokenExpiration.Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		return domain.Token{}, err
	}

	// Phát hành Refresh Token (hết hạn trong 7 ngày)
	refreshTokenExpiration := time.Now().Add(7 * 24 * time.Hour)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    refreshTokenExpiration.Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(jwtKey))
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{
		AccessToken:      accessTokenString,
		RefreshToken:     refreshTokenString,
	}, nil
}

func (a *authUseCase) Login(email, password string) (domain.Token, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return domain.Token{}, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.Token{}, errors.New("invalid credentials222")
	}

	return a.issueTokens(strconv.Itoa(user.ID))
}

func (a *authUseCase) Register(user *domain.SysUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	// return a.userRepo.Create(user)
	_, errCreate := a.userRepo.Create(user)
	if errCreate!= nil {
        return errCreate
    }
	return nil
}

func (a *authUseCase) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
}


func (a *authUseCase) RefreshToken(refreshTokenString string) (domain.Token, error) {
	// Parse và xác thực Refresh Token
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return domain.Token{}, errors.New("invalid refresh token")
	}

	// Lấy thông tin userID từ claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return domain.Token{}, errors.New("invalid refresh token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return domain.Token{}, errors.New("invalid refresh token payload")
	}

	// Phát hành cặp token mới cho user
	return a.issueTokens(userID)
}