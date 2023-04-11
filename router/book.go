package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

func BookRouter(app *fiber.App) {
	router := app.Group("/api")
	router.Get("/books", middleware.CheckUserApiKey, bookcontroller.FindAll)
	router.Get("/books/:id", bookcontroller.FindById)
	router.Post("/books", bookcontroller.Create)
	router.Put("/books/:id", bookcontroller.Update)
	router.Delete("/books/:id", bookcontroller.Delete)
}
