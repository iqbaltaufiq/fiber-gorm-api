package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.CreateUserRequest) fiber.Map
	Login(ctx context.Context, request web.LoginUserRequest) fiber.Map
	GenerateApiKey(ctx context.Context, username string) fiber.Map
}
