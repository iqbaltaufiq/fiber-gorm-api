package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

// constructor
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	if errCreate := tx.Create(&user).Error; errCreate != nil {
		return domain.User{}, errCreate
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindUserById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	var user domain.User
	errFind := tx.Where("id = ?", userId).First(&user).Error
	if errFind != nil {
		return domain.User{}, errFind
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindUserByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.User, error) {
	var user domain.User

	errFind := tx.Where("username = ?", username).First(&user).Error
	if errFind != nil {
		// panic(errFind.Error)
		return domain.User{}, errFind
	}

	return user, nil
}

func (r *UserRepositoryImpl) CreateApiKey(ctx context.Context, tx *gorm.DB, key domain.ApiKey) (domain.ApiKey, error) {
	if errCreate := tx.Create(&key).Error; errCreate != nil {
		return domain.ApiKey{}, errCreate
	}

	return key, nil
}
