package bookcontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
)

// Create creates new book entry into the database
func Create(c *fiber.Ctx) error {
	var requestBody domain.Book

	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed to parse request body.",
			},
		})
	}

	err = model.DB.Create(&requestBody).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed to insert an entry into database.",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   requestBody,
	})
}

// FindAll retrieves all of the books in the database (no limit)
func FindAll(c *fiber.Ctx) error {
	var books []domain.Book
	err := model.DB.Find(&books).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "Entry not found.",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   books,
	})
}

// FindById retrieve a book from database
func FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	var book domain.Book

	err := model.DB.First(&book, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": "Entry not found.",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

// Update modify a book entry in database
func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var book domain.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed to parse request body.",
			},
		})
	}

	err = model.DB.Where("id = ?", id).Updates(&book).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": "Failed to update an entry.",
			},
		})
	}

	idInt, _ := strconv.Atoi(id)
	book.Id = int64(idInt)

	return c.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   book,
	})

}

// Delete delets a book from database
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book domain.Book

	err := model.DB.Delete(&book, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Bad Request",
			Data: fiber.Map{
				"message": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "An entry has been deleted.",
		},
	})
}
