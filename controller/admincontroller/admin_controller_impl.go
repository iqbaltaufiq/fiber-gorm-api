package admincontroller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
)

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(AdminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: AdminService,
	}
}

func (c *AdminControllerImpl) Create(ctx *fiber.Ctx) error {
	var adminCreate web.CreateAdminRequest
	errParse := ctx.BodyParser(&adminCreate)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	res := c.AdminService.Create(ctx.Context(), adminCreate)
	if res["error"] == true {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": res["err_msg"],
			},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   res["result"],
	})
}

func (c *AdminControllerImpl) Login(ctx *fiber.Ctx) error {
	var credential web.LoginAdminRequest
	errParse := ctx.BodyParser(&credential)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	res := c.AdminService.Login(ctx.Context(), credential)
	if res["error"] != false {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Login failed. Please check again.",
			},
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "a_auth",
		Value:   "admin login",
		Path:    "/",
		Domain:  "localhost",
		Expires: time.Now().Add(30 * 24 * time.Hour),
	})

	return ctx.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "Login success.",
		},
	})
}
