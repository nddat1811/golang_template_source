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
	ExportUsersToTemplate() (*bytes.Buffer, error)
	EditFullName(id int, fullName string) (*domain.SysUser, error)
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
	sheet := f.GetSheetName(0)

	// Write header
	headers := []string{"ID", "Name", "Email", "CreatedAt", "UpdatedAt"}
	for col, header := range headers {
  	cell := string(rune('A'+col)) + "1"
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


func (u *userUseCase) ExportUsersToTemplate() (*bytes.Buffer, error) {
	// Open the template Excel file
	f, err := excelize.OpenFile("templates/template_user.xlsx")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get all users
	users, err := u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	sheet := f.GetSheetName(0)
	rowIndex := 2 // Start writing from row 2, assuming row 1 is header

	for i, user := range users {
		f.SetCellValue(sheet, "A"+strconv.Itoa(rowIndex+i), i+1)       // STT
		f.SetCellValue(sheet, "B"+strconv.Itoa(rowIndex+i), user.Name) // Name
		f.SetCellValue(sheet, "C"+strconv.Itoa(rowIndex+i), user.Email) // Email
	}

	// Write to buffer
	buffer := new(bytes.Buffer)
	if err := f.Write(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (u *userUseCase) EditFullName(id int, fullName string) (*domain.SysUser, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = fullName

	updatedUser, err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}