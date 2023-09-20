package handler

import (
	"Si-KP/database"
	"Si-KP/response"
	"Si-KP/services"
	"Si-KP/table"
	"errors"
	"log"

	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Get Books
// @Description Get a list of books with pagination and filtering options
// @Tags Books
// @Security    ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of results per page"
// @Param nama query string false "Nama"
// @Success 200 {object} response.GetBooks
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /books [get]
func GetBooks(c *fiber.Ctx) error {
	DB := database.DB

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	nama := c.Query("nama")
	offset := (page - 1) * limit
	if len(nama) < 3 && nama != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "nama must be at least 3 characters long",
		})
	}
	var (
		total_rows int64
		data       []table.Books
		get        []response.GetBooks
	)

	query := DB.Model(&data).
		Where("nama LIKE ?", "%"+nama+"%").
		// Where("books.deleted_at IS NULL").
		Order("id desc")

	query.Count(&total_rows)
	query.Offset(offset).Limit(limit)
	query.Scan(&get)
	if len(get) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No data found.",
		})
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Get Data Books Succes",
		"data":    get,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total_rows,
		},
	})

}

// @Summary Create Books
// @Description Create a Books
// @Tags Books
// @Accept json
// @Produce json
// @Param request body response.FormBooks true "Books data"
// @Success 200 {object} response.FormBooks
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /books [post]
// @Security    ApiKeyAuth
func CreateBooks(c *fiber.Ctx) error {

	DB := database.DB

	var (
		request = new(response.FormBooks)
		books   table.Books
	)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	// exists, err := services.IsUsernameExists(request.Badge)
	// if err != nil {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": err,
	// 	})
	// }
	// if exists {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Badge sudah Terdaftar",
	// 	})
	// }
	books.Title = request.Title
	books.Nama = request.Nama
	books.Judul = request.Judul
	books.Penulis = request.Penulis
	books.Penerbit = request.Penerbit
	books.CreatedAt = services.TimeNow()
	books.Created_by = services.GetUserLogin(c)

	// check
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := DB.Create(&books).Error; err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Create Data Books",
		"request": request,
	})

}

// @Summary Soft Delete Books
// @Description Soft delete a Books by ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Books ID"
// @Success 200 {string} response.GetBooks
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /books/{id} [delete]
// @Security    ApiKeyAuth
func SoftDeleteBooks(c *fiber.Ctx) error {
	userID := c.Params("id")

	var books table.Books
	errCheckUser := database.DB.Debug().First(&books, "id = ?", userID).Error
	if errCheckUser != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found!",
		})
	}
	books.DeletedAt.Time = services.TimeNow()
	books.DeletedAt.Valid = true

	if err := database.DB.Save(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to soft delete user"})
	}
	return c.JSON(fiber.Map{
		"message": "Delete Data Succes!",
		"data":    books,
	})
}

// @Summary Update Books
// @Description Update a new Books
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Books ID"
// @Param request body response.FormBooks true "Update User data"
// @Success 200 {object} response.FormBooks
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /books/{id} [put]
// @Security    ApiKeyAuth
func UpdateBooks(c *fiber.Ctx) error {
	BooksID := c.Params("id")

	var books table.Books
	err := database.DB.First(&books, BooksID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "books not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve books",
		})
	}

	var request response.FormBooks
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	books.Title = request.Title
	books.Nama = request.Nama
	books.Judul = request.Judul
	books.Penulis = request.Penulis
	books.Penerbit = request.Penerbit

	if err := database.DB.Save(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Books",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update Books success",
		"data":    request,
	})
}
