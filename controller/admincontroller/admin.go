package admincontroller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *fiber.Ctx) error {
	var payload domain.Admin

	errParse := c.BodyParser(&payload)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	payload.Role = "admin"

	// hash the password before inserted into DB
	hash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(hash)

	errCreate := model.DB.Create(&payload).Error
	if errCreate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed creating new admin.",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data: fiber.Map{
			"message": "A new admin has been created successfully.",
			"admin":   payload,
		},
	})
}

func Login(c *fiber.Ctx) error {

	var credential domain.Admin
	errParse := c.BodyParser(&credential)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	var adminInDB domain.Admin

	errFind := model.DB.Where("username = ?", credential.Username).First(&adminInDB).Error
	if errFind != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "No entry found.",
			},
		})
	}

	// password checking
	errHash := bcrypt.CompareHashAndPassword([]byte(adminInDB.Password), []byte(credential.Password))
	if errHash != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Wrong password.",
			},
		})
	}

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "auth",
		Value:   "admin login",
		Path:    "/",
		Domain:  "localhost",
		Expires: time.Now().Add(30 * 24 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "Login success",
		},
	})
}
