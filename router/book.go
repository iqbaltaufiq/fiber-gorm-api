package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

// route for user to query book(s).
// user with role 'user' can only
// retrieve book(s).
func BookRouter(app *fiber.App, middleware *middleware.AuthMiddleware, bookcontroller bookcontroller.BookController) {
	router := app.Group("/api")
	router.Get("/books", middleware.CheckUserApiKey, bookcontroller.FindAll)
	router.Get("/books/:id", middleware.CheckUserApiKey, bookcontroller.FindById)
}
