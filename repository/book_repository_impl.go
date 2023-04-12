package repository

import (
	"context"

	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	errCreate := tx.Create(&book).Error
	if errCreate != nil {
		return domain.Book{}, errCreate
	}

	return book, nil
}

func (r *BookRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	errUpdate := tx.Where("id = ?", book.Id).Updates(&book).Error
	if errUpdate != nil {
		return domain.Book{}, errUpdate
	}

	return book, nil
}

func (r *BookRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, bookId int) error {
	errDelete := tx.Delete(domain.Book{}, bookId).Error
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (r *BookRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, bookId int) (domain.Book, error) {
	var book domain.Book
	errFind := tx.First(&book, bookId).Error
	if errFind != nil {
		return domain.Book{}, errFind
	}

	return book, nil
}

func (r *BookRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Book, error) {
	var books []domain.Book
	errFind := tx.Find(&books).Error
	if errFind != nil {
		return []domain.Book{}, errFind
	}

	return books, nil
}
