package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/admincontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/bookcontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/controller/usercontroller"
	"github.com/iqbaltaufiq/go-fiber-restapi/middleware"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
	"gorm.io/gorm"
)

func NewRouter(app *fiber.App, middleware *middleware.AuthMiddleware, db *gorm.DB, validate *validator.Validate) {
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := bookcontroller.NewBookController(bookService)

	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository, db, validate)
	adminController := admincontroller.NewAdminController(adminService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := usercontroller.NewUserController(userService)

	// put routes here
	BookRouter(app, middleware, bookController)
	AdminRouter(app, middleware, adminController, bookController)
	UserRouter(app, userController)
}
