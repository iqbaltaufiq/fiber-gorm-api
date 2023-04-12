package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

type AdminService interface {
	Create(ctx context.Context, request web.CreateAdminRequest) fiber.Map
	Login(ctx context.Context, credential web.LoginAdminRequest) fiber.Map
}
