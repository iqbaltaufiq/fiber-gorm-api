package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (r *AdminRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, admin domain.Admin) (domain.Admin, error) {
	if errCreate := tx.Create(&admin).Error; errCreate != nil {
		return domain.Admin{}, errCreate
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.Admin, error) {
	var admin domain.Admin
	errFind := tx.Where("username = ?", username).First(&admin).Error
	if errFind != nil {
		return domain.Admin{}, errFind
	}

	return admin, nil
}
