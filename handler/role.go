package handler

import (
	"Si-KP/database"
	"Si-KP/table"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// // @Summary Create Role
// // @Description Create a new user
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Param request body response.RoleCreateRequest true "User data"
// // @Success 200 {object} response.RoleCreateRequest
// // @Failure 400 {object} response.ErrorResponse
// // @Failure 404 {object} response.ErrorResponse
// // @Router /role-create [post]
// // @Security    ApiKeyAuth
// func CreateRole(c *fiber.Ctx) error {
// 	DB := database.DB
// 	var (
// 		request = new(response.RoleCreateRequest)
// 		role    table.Role
// 	)
// 	if err := c.BodyParser(&request); err != nil {
// 		log.Println(err)
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": err,
// 		})
// 	}

// 	role.Role = request.Role
// 	role.Created_at = services.TimeNow()
// 	role.Created_by = services.GetUserLogin(c)

// 	validate := validator.New()
// 	if err := validate.Struct(request); err != nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	// Check if a role with the same name already exists in the database
// 	existingRole := table.Role{}
// 	if err := DB.Where("role = ?", role.Role).First(&existingRole).Error; err == nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Role already exists",
// 		})
// 	}

// 	if err := DB.Create(&role).Error; err != nil {
// 		log.Println(err)
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": err,
// 		})
// 	}

// 	c.Status(200)
// 	return c.JSON(fiber.Map{
// 		"message": "Succes Create Data",
// 		"request": request,
// 	})

// }

// @Summary Get Role
// @Description Get a list of Role
// @Tags Role
// @Security    ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of results per page"
// @Success 200 {object} response.RoleResponseGet
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /role [get]
func GetRole(c *fiber.Ctx) error {

	DB := database.DB

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var (
		total_rows int64
		data       []table.Role
	)

	query := DB.Table("roles").
		Where("deleted_at IS NULL").
		Order("id desc")

	query.Count(&total_rows)
	query.Offset(offset).Limit(limit)
	query.Find(&data)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Get Data Role",
		"data":    data,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total_rows,
		},
	})

}

// // @Summary Update Role
// // @Description Update Role
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Param id path int true "Role ID"
// // @Param request body response.RoleUpdateRequest true "Update Role data"
// // @Success 200 {object} response.RoleUpdateRequest
// // @Failure 400 {object} response.ErrorResponse
// // @Failure 404 {object} response.ErrorResponse
// // @Router /role/{id} [put]
// // @Security    ApiKeyAuth
// func UpdateRole(c *fiber.Ctx) error {
// 	userID := c.Params("id")
// 	DB := database.DB

// 	var role table.Role
// 	err := DB.First(&role, userID).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"message": "Role not found",
// 			})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to retrieve Role",
// 		})
// 	}
// 	var request response.RoleUpdateRequest
// 	if err := c.BodyParser(&request); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 	}
// 	validate := validator.New()
// 	if err := validate.Struct(request); err != nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	// Check if a role with the same name already exists in the database
// 	var existingRole table.Role
// 	if err := DB.Where("role = ?", request.Role).First(&existingRole).Error; err == nil {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "Role already exists",
// 		})
// 	}

// 	role.Role = request.Role

// 	if err := database.DB.Save(&role).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to update Role",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Update Role success",
// 		"data":    request.Role,
// 	})
// }

// // @Summary Soft Delete Role
// // @Description Soft delete Role
// // @Tags Role
// // @Accept json
// // @Produce json
// // @Param id path int true "Role ID"
// // @Success 200 {string} response.RoleResponseGet
// // @Failure 400 {object} response.ErrorResponse
// // @Failure 404 {object} response.ErrorResponse
// // @Router /role/{id} [delete]
// // @Security    ApiKeyAuth
// func SoftDeleteRole(c *fiber.Ctx) error {
// 	userID := c.Params("id")

// 	var Role table.Role
// 	errCheckUser := database.DB.Debug().First(&Role, "id = ?", userID).Error
// 	if errCheckUser != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"message": "Role not found!",
// 		})
// 	}
// 	Role.Deleted_by = services.GetUserLogin(c)
// 	Role.DeletedAt.Time = services.TimeNow()
// 	Role.DeletedAt.Valid = true

// 	if err := database.DB.Save(&Role).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to soft delete Role"})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Delete Data Succes!",
// 		"data":    Role,
// 	})
// }
