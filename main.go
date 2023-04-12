package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"github.com/iqbaltaufiq/go-fiber-restapi/router"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := model.SetupDatabase()
	validate := validator.New()
	model.RunMigration()

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := bookcontroller.NewBookController(bookService)

	app := fiber.New()

	middleware.SetupMiddleware(app)

	// router here
	router.BookRouter(app, bookController)
	router.AdminRouter(app, bookController)
	router.UserRouter(app)

	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
