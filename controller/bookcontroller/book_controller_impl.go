package bookcontroller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/service"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(BookService service.BookService) BookController {
	return &BookControllerImpl{BookService: BookService}
}

func (c *BookControllerImpl) Create(ctx *fiber.Ctx) error {
	var createBookRequest web.CreateBookRequest
	errParse := ctx.BodyParser(&createBookRequest)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed to parse request body.",
			},
		})
	}

	res := c.BookService.Create(ctx.Context(), createBookRequest)
	if res["error"] == true {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed to insert an entry into database.",
			},
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.StdResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   res["result"],
	})
}

func (c *BookControllerImpl) Update(ctx *fiber.Ctx) error {
	var updateBookRequest web.UpdateBookRequest
	id := ctx.Params("id")

	errParse := ctx.BodyParser(&updateBookRequest)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed to parse request body.",
			},
		})
	}

	idInt, _ := strconv.Atoi(id)
	updateBookRequest.Id = int64(idInt)

	res := c.BookService.Update(ctx.Context(), updateBookRequest)
	if res["error"] == true {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.StdResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: fiber.Map{
				"message": "Failed to update a book in database.",
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res["result"],
	})
}

func (c *BookControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	bookId, _ := strconv.Atoi(id)

	res := c.BookService.Delete(ctx.Context(), bookId)
	if res["error"] == true {
		return ctx.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data: fiber.Map{
				"message": res["err_msg"],
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.StdResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"message": "A book has been deleted.",
		},
	})
}

func (c *BookControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	res := c.BookService.FindById(ctx.Context(), id)
	if res["error"] == true {
		return ctx.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
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

func (c *BookControllerImpl) FindAll(ctx *fiber.Ctx) error {
	res := c.BookService.FindAll(ctx.Context())
	if res["error"] == true {
		return ctx.Status(fiber.StatusNotFound).JSON(web.StdResponse{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
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
