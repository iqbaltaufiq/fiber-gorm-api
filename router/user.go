package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/usercontroller"
)

// router for handling user's credential stuffs
// like creating new user
// logging in and generating new API key.
func UserRouter(c *fiber.App) {
	router := c.Group("/api/user")
	router.Post("/register", usercontroller.Register)
	router.Post("/login", usercontroller.Login)
	router.Post("/newkey", usercontroller.GenerateApiKey)
}
