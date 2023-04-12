package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/admincontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
	"gorm.io/gorm"
)

func NewRouter(app *fiber.App, db *gorm.DB, validate *validator.Validate) {
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := bookcontroller.NewBookController(bookService)

	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository, db, validate)
	adminController := admincontroller.NewAdminController(adminService)

	// put routes here
	BookRouter(app, bookController)
	AdminRouter(app, adminController, bookController)
	UserRouter(app)
}
