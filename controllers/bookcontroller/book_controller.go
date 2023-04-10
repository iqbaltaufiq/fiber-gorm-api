package bookcontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/helper"
	"github.com/iqbaltaufiq/go-fiber-restapi/models"
)

type RespMessage map[string]string

// Create creates new book entry into the database
func Create(c *fiber.Ctx) error {
	var requestBody models.Book

	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: RespMessage{
				"message": "Failed to parse request body.",
			},
		})
	}

	err = models.DB.Create(&requestBody).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: RespMessage{
				"message": "Failed to insert an entry into database.",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(helper.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   requestBody,
	})
}

func FindAll(c *fiber.Ctx) error {
	var books []models.Book

	err := models.DB.Find(&books).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: RespMessage{
				"message": "Entry not found.",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(helper.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   books,
	})
}

func FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	err := models.DB.First(&book, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: RespMessage{
				"message": "Entry not found.",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(helper.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: RespMessage{
				"message": "Failed to parse request body.",
			},
		})
	}

	err = models.DB.Where("id = ?", id).Updates(&book).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: RespMessage{
				"message": "Failed to update an entry.",
			},
		})
	}

	idInt, _ := strconv.Atoi(id)
	book.Id = int64(idInt)

	return c.Status(fiber.StatusOK).JSON(helper.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   book,
	})

}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	err := models.DB.Delete(&book, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Bad Request",
			Data: RespMessage{
				"message": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(helper.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: RespMessage{
			"message": "An entry has been deleted.",
		},
	})
}
