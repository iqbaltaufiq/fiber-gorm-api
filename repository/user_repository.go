package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	FindUserById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error)
	FindUserByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.User, error)
	CreateApiKey(ctx context.Context, tx *gorm.DB, key domain.ApiKey) (domain.ApiKey, error)
}
