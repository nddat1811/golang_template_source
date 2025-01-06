package repository

import (
	"gorm.io/gorm"
	"golang_template_source/domain"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetByID(id int) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
