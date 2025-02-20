//go:generate mockgen -source user.go -destination ../mock/mock_repository/user_gen.go
package repository

import (
	"golang_template_source/internal/domain"
	"golang_template_source/internal/domain/convert"
	"golang_template_source/internal/domain/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll() ([]*domain.SysUser, error)
	GetByID(id int) (*domain.SysUser, error)
	FindByEmail(email string) (*domain.SysUser, error)
	Create(user *domain.SysUser) (id int, err error)
	UpdateUser(user *domain.SysUser) (*domain.SysUser, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetAll() ([]*domain.SysUser, error) {
	var users []*domain.SysUser
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetByID(id int) (*domain.SysUser, error) {
	var user entity.SysUser
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Chuyển đổi từ entity sang domain
	userDomain := convert.ConvertEntityToDomain(&user)
	return userDomain, nil
	// return &user, nil
}


func (u *userRepository) FindByEmail(email string) (*domain.SysUser, error) {
	var user entity.SysUser
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Chuyển đổi từ entity sang domain
	userDomain := convert.ConvertEntityToDomain(&user)
	return userDomain, nil
}

func (u *userRepository) Create(user *domain.SysUser) (id int, err error) {

	err = u.db.Create(&user).Error

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (u *userRepository) UpdateUser(user *domain.SysUser) (*domain.SysUser, error) {
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}