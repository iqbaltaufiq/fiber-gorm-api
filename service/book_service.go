package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

type BookService interface {
	Create(ctx context.Context, request web.CreateBookRequest) fiber.Map
	Update(ctx context.Context, request web.UpdateBookRequest) fiber.Map
	Delete(ctx context.Context, bookId int) fiber.Map
	FindById(ctx context.Context, bookId int) fiber.Map
	FindAll(ctx context.Context) fiber.Map
}
