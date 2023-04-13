package usercontroller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// constructor
func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{UserService: UserService}
}

func (c *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	var request web.CreateUserRequest
	errParse := ctx.BodyParser(&request)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	res := c.UserService.Register(ctx.Context(), request)
	if res["error"] != false {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": res["err_msg"],
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res["result"],
	})
}

func (c *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	var request web.LoginUserRequest
	errParse := ctx.BodyParser(&request)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	res := c.UserService.Login(ctx.Context(), request)
	if res["error"] != false {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": res["err_msg"],
			},
		})
	}

	cookieValue := fmt.Sprintf(`{"u":"%s","r":"user"}`, request.Username)

	ctx.Cookie(&fiber.Cookie{
		Name:    "u_auth",
		Path:    "/",
		Domain:  "localhost",
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Value:   cookieValue,
	})

	return ctx.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "Login success.",
		},
	})
}

func (c *UserControllerImpl) GenerateApiKey(ctx *fiber.Ctx) error {
	var cookie fiber.Map
	cookieJSON := ctx.Cookies("u_auth")
	errParse := json.Unmarshal([]byte(cookieJSON), &cookie)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": "Failed parsing user cookie.",
			},
		})
	}

	username, _ := cookie["u"].(string)

	res := c.UserService.GenerateApiKey(ctx.Context(), username)
	if res["error"] != false {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.ErrBadRequest.Code,
			Status: fiber.ErrBadRequest.Message,
			Data: fiber.Map{
				"message": "Failed generating new api key.",
			},
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
		Code:   fiber.ErrBadRequest.Code,
		Status: fiber.ErrBadRequest.Message,
		Data:   res["result"],
	})
}
