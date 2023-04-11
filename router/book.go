package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
)

// route for user to query book(s).
// user with role 'user' can only
// retrieve book(s).
func BookRouter(app *fiber.App) {
	router := app.Group("/api", middleware.CheckUserApiKey)
	router.Get("/books", bookcontroller.FindAll)
	router.Get("/books/:id", bookcontroller.FindById)
}
