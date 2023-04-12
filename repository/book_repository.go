package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	Save(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	Update(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	Delete(ctx context.Context, tx *gorm.DB, bookId int) error
	FindById(ctx context.Context, tx *gorm.DB, bookId int) (domain.Book, error)
	FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Book, error)
}
