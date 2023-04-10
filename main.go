package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controllers/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/models"
)

func main() {
	// create a db connection
	// and run migration
	models.NewDBConnection()

	app := fiber.New()

	api := app.Group("/api")
	api.Get("/books", bookcontroller.FindAll)
	api.Get("/books/:id", bookcontroller.FindById)
	api.Post("/books", bookcontroller.Create)
	api.Put("/books/:id", bookcontroller.Update)
	api.Delete("/books/:id", bookcontroller.Delete)

	app.Listen(":3000")
}
