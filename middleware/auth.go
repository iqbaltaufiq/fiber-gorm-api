package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

var CheckUserAuth = func(c *fiber.Ctx) error {
	if c.GetReqHeaders()["X-Api-Key"] != "SECRET" {
		return errors.New("Unauthorized")
	}

	return c.Next()
}

var CheckAdminAuth = func(c *fiber.Ctx) error {
	fmt.Printf("Cookie: %s\n\n", c.Cookies("auth"))
	if c.Cookies("auth") != "admin login" {
		return c.Status(fiber.StatusUnauthorized).JSON(web.StdResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data: fiber.Map{
				"message": "Admin only. Please login first.",
			},
		})
	}

	return c.Next()
}
