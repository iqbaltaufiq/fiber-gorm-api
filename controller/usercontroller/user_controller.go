package usercontroller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	GenerateApiKey(ctx *fiber.Ctx) error
}
