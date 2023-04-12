package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Save(ctx context.Context, tx *gorm.DB, admin domain.Admin) (domain.Admin, error)
	FindByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.Admin, error)
}
