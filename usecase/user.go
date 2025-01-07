package usecase

import (
	"bytes"
	"github.com/xuri/excelize/v2"
	"golang_template_source/domain"
	"golang_template_source/repository"
	"strconv"
)

type UserUseCase interface {
	FindByEmail(email string) (*domain.SysUser, error)
	Create(user *domain.SysUser) error
	GetAllUsers() ([]*domain.SysUser, error)
	GetUserByID(id int) (*domain.SysUser, error)
	ExportUsersToExcel() (*bytes.Buffer, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) GetAllUsers() ([]*domain.SysUser, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) GetUserByID(id int) (*domain.SysUser, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) FindByEmail(email string) (*domain.SysUser, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userUseCase) Create(user *domain.SysUser) error {
	_, err := u.userRepo.Create(user)
	if err!= nil {
        return err
    }
	return nil
}

func (u *userUseCase) ExportUsersToExcel() (*bytes.Buffer, error) {
	users, err := u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Users"
	f.NewSheet(sheet)

	// Write header
	headers := []string{"ID", "Name", "Email", "CreatedAt", "UpdatedAt"}
	for col, header := range headers {
		cell := string('A'+col) + "1"
		f.SetCellValue(sheet, cell, header)
	}

	// Write user data
	for row, user := range users {
		f.SetCellValue(sheet, "A"+strconv.Itoa(row+2), user.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row+2), user.Name)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row+2), user.Email)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row+2), user.CreatedAt)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row+2), user.UpdatedAt)
	}

	buffer := new(bytes.Buffer)
	if err := f.Write(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}
