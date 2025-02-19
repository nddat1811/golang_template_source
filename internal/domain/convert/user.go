package convert

import (
	"golang_template_source/internal/domain"
	"golang_template_source/internal/domain/entity"
)

// ConvertDomainToEntity chuyển đổi từ domain model sang entity model và trả về con trỏ
func ConvertDomainToEntity(user *domain.SysUser) *entity.SysUser {
	if user == nil {
		return nil
	}

	return &entity.SysUser{
		ID:        user.ID,
		Email:     user.Email,
		RandomID:  user.RandomID,
		Password:  user.Password,
		Phone:     user.Phone,
		Name:      user.Name,
		Status:    user.Status,
		Identity:  user.Identity,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
	}
}

// ConvertEntityToDomain chuyển đổi từ entity model sang domain model và trả về con trỏ
func ConvertEntityToDomain(entity *entity.SysUser) *domain.SysUser {
	if entity == nil {
		return nil
	}

	return &domain.SysUser{
		ID:        entity.ID,
		Email:     entity.Email,
		RandomID:  entity.RandomID,
		Password:  entity.Password,
		Phone:     entity.Phone,
		Name:      entity.Name,
		Status:    entity.Status,
		Identity:  entity.Identity,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
	}
}
