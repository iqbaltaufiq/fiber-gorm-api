package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

var CheckUserApiKey = func(c *fiber.Ctx) error {
	var keyInDB domain.ApiKey
	apiKey := c.GetReqHeaders()["X-Api-Key"]
	errFind := model.DB.Where("api_key = ?", apiKey).First(&keyInDB).Error
	if errFind != nil || keyInDB.Expires.Unix() < time.Now().Unix() {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "API key doesn't exist.",
			},
		})
	}

	return c.Next()
}

var CheckUserAuth = func(c *fiber.Ctx) error {
	var cookie fiber.Map
	json.Unmarshal([]byte(c.Cookies("u_auth")), &cookie)

	if cookie["r"] != "user" {
		return c.Status(fiber.StatusUnauthorized).JSON(web.StdResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data: fiber.Map{
				"message": "Please login first before creating a new API key.",
			},
		})
	}

	return c.Next()
}

var CheckAdminAuth = func(c *fiber.Ctx) error {
	fmt.Printf("Cookie: %s\n\n", c.Cookies("auth"))
	if c.Cookies("auth") != "admin login" {
		return c.Status(fiber.StatusUnauthorized).JSON(web.StdResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data: fiber.Map{
				"message": "Admin only. Please login first.",
			},
		})
	}

	return c.Next()
}
