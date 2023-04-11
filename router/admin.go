package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/admincontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

func AdminRouter(app *fiber.App) {
	router := app.Group("/api/admin")
	router.Get("/books", middleware.CheckAdminAuth, bookcontroller.FindAll)
	router.Get("/books/:id", middleware.CheckAdminAuth, bookcontroller.FindById)
	router.Post("/books", middleware.CheckAdminAuth, bookcontroller.Create)
	router.Put("/books/:id", middleware.CheckAdminAuth, bookcontroller.Update)
	router.Delete("/books/:id", middleware.CheckAdminAuth, bookcontroller.Delete)

	router.Post("/new", admincontroller.Create)
	router.Post("/login", admincontroller.Login)
}
