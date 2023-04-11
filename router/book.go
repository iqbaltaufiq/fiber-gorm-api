package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

func BookRouter(app *fiber.App) {
	router := app.Group("/api", middleware.CheckUserApiKey)
	router.Get("/books", bookcontroller.FindAll)
	router.Get("/books/:id", bookcontroller.FindById)
}
