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

// @Summary Get Users
// @Description Get a list of users with pagination and filtering options
// @Tags Users
// @Security    ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of results per page"
// @Param badge query string false "User badge"
// @Success 200 {object} response.GetUsers
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	DB := database.DB

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	badge := c.Query("badge")
	offset := (page - 1) * limit
	if len(badge) < 3 && badge != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Badge must be at least 3 characters long",
		})
	}
	var (
		total_rows int64
		data       []table.Users
		get        []response.GetUsers
	)

	query := DB.Model(&data).
		Select("users.*", "roles.role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("badge LIKE ?", "%"+badge+"%").
		Where("users.deleted_at IS NULL").
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
		"message": "Get Data Users Succes",
		"data":    get,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total_rows,
		},
	})

}

// @Summary Create User
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body response.FormUsers true "User data"
// @Success 200 {object} response.FormUsers
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /users [post]
// @Security    ApiKeyAuth
func CreateUsers(c *fiber.Ctx) error {

	DB := database.DB

	var (
		request = new(response.FormUsers)
		users   table.Users
	)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	exists, err := services.IsUsernameExists(request.Badge)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	if exists {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Badge sudah Terdaftar",
		})
	}
	users.Badge = request.Badge
	users.Nama = request.Nama
	users.Email = request.Email
	users.Password, _ = services.HashPassword(request.Password)
	users.Created_at = services.TimeNow()
	users.Created_by = services.GetUserLogin(c)

	// Check if the provided Role ID exists in the database
	role := table.Role{}
	if err := DB.First(&role, request.Role_Id).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Role ID",
		})
	}
	users.Role_Id = request.Role_Id
	// check
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := DB.Create(&users).Error; err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Succes Create Data User",
		"request": request,
	})

}

// @Summary Soft Delete User
// @Description Soft delete a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} response.GetUsers
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /users/{id} [delete]
// @Security    ApiKeyAuth
func SoftDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user table.Users
	errCheckUser := database.DB.Debug().First(&user, "id = ?", userID).Error
	if errCheckUser != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found!",
		})
	}
	user.Deleted_by = services.GetUserLogin(c)
	user.DeletedAt.Time = services.TimeNow()
	user.DeletedAt.Valid = true

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to soft delete user"})
	}
	return c.JSON(fiber.Map{
		"message": "Delete Data Succes!",
		"data":    user,
	})
}

// @Summary Update User
// @Description Update a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body response.FormUsers true "Update User data"
// @Success 200 {object} response.FormUsers
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /users/{id} [put]
// @Security    ApiKeyAuth
func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user table.Users
	err := database.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
		})
	}

	var request response.FormUsers
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
	exists, err := services.IsUsernameExists(request.Badge)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to check username existence",
		})
	}
	if exists && request.Badge != user.Badge {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already registered",
		})
	}

	user.Badge = request.Badge
	user.Nama = request.Nama
	user.Email = request.Email
	user.Role_Id = request.Role_Id
	user.Updated_at = services.TimeNow()
	user.Updated_by = services.GetUserLogin(c)

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Update user success",
		"data":    request,
	})
}
