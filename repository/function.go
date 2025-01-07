package repository

import (
	"golang_template_source/domain"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

// SysFunctionRepository provides access to function data
type sysFunctionRepository struct {
	db *gorm.DB
}

type SysFunctionRepository interface {
	CheckAndReturnOriginalPath(urlPath string, pathPrefix string) (string, error)
	IsAuthentication(userID int, url string) (bool, error)
}

// NewSysFunctionRepository creates a new repository instance
func NewSysFunctionRepository(db *gorm.DB) SysFunctionRepository {
	return &sysFunctionRepository{db: db}
}

func (r *sysFunctionRepository) CheckAndReturnOriginalPath(urlPath string, pathPrefix string) (string, error) {
	urlPath2 := strings.Replace(urlPath, "/"+pathPrefix, "", 1)
	if urlPath2 == "" {
		return "x", nil
	}
	var functions []domain.SysFunction
	result := r.db.Where("regex IS NOT NULL").Find(&functions)
	if result.Error != nil {
		return urlPath, result.Error
	}

	for _, function := range functions {
		if function.Regex == nil {
			continue
		}
		pattern, err := regexp.Compile(*function.Regex)
		if err != nil {
			continue
		}
		if pattern.MatchString(urlPath) {
			return function.Path, nil
		}
	}

	return urlPath, nil
}

func (r *sysFunctionRepository) IsAuthentication(userID int, url string) (bool, error) {
	var count int64
	result := r.db.Table("SYS_USER").
		Select("COUNT(*)").
		Joins("JOIN SYS_USER_ROLE ON SYS_USER.id = SYS_USER_ROLE.user_id").
		Joins("JOIN SYS_ROLE ON SYS_USER_ROLE.role_id = SYS_ROLE.id").
		Joins("JOIN SYS_ROLE_FUNCTION ON SYS_ROLE.id = SYS_ROLE_FUNCTION.role_id").
		Joins("JOIN SYS_FUNCTION ON SYS_ROLE_FUNCTION.function_id = SYS_FUNCTION.id").
		Where("SYS_USER.id = ? AND SYS_FUNCTION.path = ?", userID, url).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
