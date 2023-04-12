package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/helper"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewBookService(BookRepository repository.BookRepository, DB *gorm.DB, Validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: BookRepository,
		DB:             DB,
		Validate:       Validate,
	}
}

func (s *BookServiceImpl) Create(ctx context.Context, request web.CreateBookRequest) fiber.Map {
	errVal := s.Validate.Struct(request)
	if errVal != nil {
		panic(errVal)
	}

	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	// mapping
	bookRequest := domain.Book{
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
		PublishDate: request.PublishDate,
	}

	book, errCreate := s.BookRepository.Save(ctx, tx, bookRequest)
	if errCreate != nil {
		return fiber.Map{
			"error":     true,
			"error_msg": errCreate.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": book,
	}
}

func (s *BookServiceImpl) Update(ctx context.Context, request web.UpdateBookRequest) fiber.Map {
	errVal := s.Validate.Struct(request)
	if errVal != nil {
		panic(errVal)
	}

	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	getBook, errFind := s.BookRepository.FindById(ctx, tx, int(request.Id))
	if errFind != nil {
		panic(errFind)
	}

	// mapping
	getBook.Title = request.Title
	getBook.Description = request.Description
	getBook.Author = request.Author
	getBook.PublishDate = request.PublishDate

	book, errUpdate := s.BookRepository.Update(ctx, tx, getBook)
	if errUpdate != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errUpdate.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": book,
	}
}

func (s *BookServiceImpl) Delete(ctx context.Context, bookId int) fiber.Map {
	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	errDelete := s.BookRepository.Delete(ctx, tx, bookId)
	if errDelete != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errDelete.Error(),
		}
	}

	return fiber.Map{
		"error": false,
	}
}

func (s *BookServiceImpl) FindById(ctx context.Context, bookId int) fiber.Map {
	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	book, errFind := s.BookRepository.FindById(ctx, tx, bookId)
	if errFind != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errFind.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": book,
	}
}

func (s *BookServiceImpl) FindAll(ctx context.Context) fiber.Map {
	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	books, errFind := s.BookRepository.FindAll(ctx, tx)
	if errFind != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errFind.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": books,
	}
}
