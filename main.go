package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	model.SetupDatabase()
	model.RunMigration()

	app := fiber.New()

	middleware.SetupMiddleware(app)

	// router here
	router.BookRouter(app)
	router.AdminRouter(app)

	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
