package repository

import (
	"golang_template_source/domain"

	"gorm.io/gorm"
)

// SysLogRepository provides access to function data
type sysLogRepository struct {
	db *gorm.DB
}

type SysLogRepository interface {
	InsertLog(log *domain.SysLog) error
}

// NewSysLogRepository creates a new repository instance
func NewSysLogRepository(db *gorm.DB) SysLogRepository {
	return &sysLogRepository{db: db}
}

func (l *sysLogRepository) InsertLog(log *domain.SysLog) error {
	return l.db.Create(log).Error
}