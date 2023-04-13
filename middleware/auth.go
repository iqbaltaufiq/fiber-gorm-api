package middleware

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	DB *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		DB: db,
	}
}

// check if the user has X-Api-Key set in their header.
// user must sent an api key to be able to query a book(s).
func (m *AuthMiddleware) CheckUserApiKey(c *fiber.Ctx) error {
	var keyInDB domain.ApiKey

	apiKey := c.GetReqHeaders()["X-Api-Key"]

	errFind := m.DB.Where("api_key = ?", apiKey).First(&keyInDB).Error
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

// check if the user has logged in
// by checking the "u_auth" cookie
func (m *AuthMiddleware) CheckUserAuth(c *fiber.Ctx) error {
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

// check if the user has logged in as admin
// to be able to access admin routes
func (m *AuthMiddleware) CheckAdminAuth(c *fiber.Ctx) error {
	if c.Cookies("a_auth") != "admin login" {
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
