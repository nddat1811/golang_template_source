//go:generate mockgen -source user.go -destination ../mock/mock_repository/user_gen.go
package repository

import (
	"golang_template_source/domain"

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
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]*domain.SysUser, error) {
	var users []*domain.SysUser
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (*domain.SysUser, error) {
	var user domain.SysUser
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *userRepository) FindByEmail(email string) (*domain.SysUser, error) {
	var user domain.SysUser
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.SysUser) (id int, err error) {

	err = r.db.Create(&user).Error

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
