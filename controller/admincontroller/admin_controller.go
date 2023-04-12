package admincontroller

import "github.com/gofiber/fiber/v2"

type AdminController interface {
	Create(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
