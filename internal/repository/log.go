package repository

import (
	"golang_template_source/internal/domain/entity"

	"gorm.io/gorm"
)

// SysLogRepository provides access to function data
type sysLogRepository struct {
	db *gorm.DB
}

type SysLogRepository interface {
	InsertLog(log *entity.SysLog) error
}

// NewSysLogRepository creates a new repository instance
func NewSysLogRepository(db *gorm.DB) SysLogRepository {
	return &sysLogRepository{db: db}
}

func (l *sysLogRepository) InsertLog(log *entity.SysLog) error {
	return l.db.Create(log).Error
}