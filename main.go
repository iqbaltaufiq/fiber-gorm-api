package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbaltaufiq/go-fiber-restapi/controllers/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/models"
)

func main() {
	// create a db connection
	// and run migration
	models.NewDBConnection()

	app := fiber.New()

	// implement CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5500",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api")
	api.Get("/books", bookcontroller.FindAll)
	api.Get("/books/:id", bookcontroller.FindById)
	api.Post("/books", bookcontroller.Create)
	api.Put("/books/:id", bookcontroller.Update)
	api.Delete("/books/:id", bookcontroller.Delete)

	app.Listen(":3000")
}
