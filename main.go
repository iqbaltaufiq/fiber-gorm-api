package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := model.OpenConnection()
	validate := validator.New()
	model.RunMigration(db)

	app := fiber.New()

	middleware.SetupMiddleware(app)
	middleware := middleware.NewAuthMiddleware(db)
	router.NewRouter(app, middleware, db, validate)

	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
