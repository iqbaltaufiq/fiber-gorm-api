package usercontroller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"golang.org/x/crypto/bcrypt"
)

// Handler for registering new user
func Register(c *fiber.Ctx) error {
	var payload domain.User
	errParse := c.BodyParser(&payload)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed parsing request body.",
			},
		})
	}

	// set role for user
	payload.Role = "user"

	// hash user password
	// before being inserted into db
	hash, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(hash)

	errCreate := model.DB.Create(&payload).Error
	if errCreate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed creating new user.",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data: fiber.Map{
			"message": "A new user has been created successfully.",
		},
	})
}

// Handler for login user
func Login(c *fiber.Ctx) error {

	var credential domain.User
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

	var userInDB domain.User

	errFind := model.DB.Where("username = ?", credential.Username).First(&userInDB).Error
	if errFind != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "No entry found.",
			},
		})
	}

	// check if the password sent by user
	// match the password in the db
	errHash := bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(credential.Password))
	if errHash != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Wrong password.",
			},
		})
	}

	cookieValue := fmt.Sprintf(`{"id":"%d","r":"user"}`, userInDB.Id)

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:    "u_auth",
		Value:   cookieValue,
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

// generate an API Key for user
// to be able to query book(s)
// the generated key is valid for 30 days
func GenerateApiKey(c *fiber.Ctx) error {
	var parseJson fiber.Map

	errParse := json.Unmarshal([]byte(c.Cookies("u_auth")), &parseJson)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed parsing cookie or cookie doesn't exist.",
			},
		})
	}

	id := parseJson["id"]

	var user domain.User

	// fetch the user from db
	errFetch := model.DB.Where("id = ?", id).First(&user).Error
	if errFetch != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "User not found.",
			},
		})
	}

	timeNow := strconv.Itoa(time.Now().Nanosecond())
	combination := timeNow + strconv.Itoa(int(user.Id))
	hash := fmt.Sprintf("%x", md5.Sum([]byte(combination)))

	// create new api key in database
	apiKey := domain.ApiKey{
		Username: user.Username,
		ApiKey:   hash,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	errCreate := model.DB.Create(&apiKey).Error
	if errCreate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed creating new API key.",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data: fiber.Map{
			"message": "A new API key has been created successfully.",
			"api_key": hash,
		},
	})
}
