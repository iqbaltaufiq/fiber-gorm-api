package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/admincontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

// route for admin activities
// admin can retrieve all books,
// retrieve a single book,
// modify and delete a book.
func AdminRouter(app *fiber.App, middleware *middleware.AuthMiddleware, admincontroller admincontroller.AdminController, bookcontroller bookcontroller.BookController) {
	router := app.Group("/api/admin")
	router.Get("/books", middleware.CheckAdminAuth, bookcontroller.FindAll)
	router.Get("/books/:id", middleware.CheckAdminAuth, bookcontroller.FindById)
	router.Post("/books", middleware.CheckAdminAuth, bookcontroller.Create)
	router.Put("/books/:id", middleware.CheckAdminAuth, bookcontroller.Update)
	router.Delete("/books/:id", middleware.CheckAdminAuth, bookcontroller.Delete)

	router.Post("/new", admincontroller.Create)
	router.Post("/login", admincontroller.Login)
}
