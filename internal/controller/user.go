package controller

import (
	"golang_template_source/internal/usecase"
	"golang_template_source/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController{
	return &UserController{userUseCase: uc}
}

// Swagger routes
// @Summary Get all users
// @Description Retrieve all users from the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.SysUser
// @Router /users [get]
func (h *UserController) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewResponse("ok", "ok"))
	// users, err := h.userUseCase.GetAllUsers()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, utils.NewResponse("Bad request", nil))
	// 	return
	// }
	// c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Retrieve a single user by ID
// @Tags users
// @Accept json
// @Produce json
// @Security Authorization
// @Param id path int true "User ID"
// @Success 200 {object} domain.SysUser
// @Router /users/{id} [get]
func (h *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse("Id invalid", nil))
		return
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewResponse("không tìm thấy user", nil))
		return
	}
	c.JSON(http.StatusOK, utils.NewResponse("ok", user))
}


// @Summary Export all users to Excel
// @Description Export all users to an Excel file and download
// @Tags users
// @Produce application/octet-stream
// @Success 200 {file} file
// @Router /users/export [get]
func (h *UserController) ExportUsersToExcel(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewResponse("ok", "ok"))
	// file, err := h.userUseCase.ExportUsersToExcel()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, utils.NewResponse("Failed to export users", nil))
	// 	return
	// }

	// c.Header("Content-Disposition", "attachment; filename=users.xlsx")
	// c.Data(http.StatusOK, "application/octet-stream", file.Bytes())
}

// @Summary Export all users to Excel from template
// @Description Export all users to an Excel file and download from template
// @Tags users
// @Produce application/octet-stream
// @Success 200 {file} file
// @Router /users/export-template [get]
func (h *UserController) ExportUsersToTemplate(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewResponse("ok", "ok"))
	// file, err := h.userUseCase.ExportUsersToTemplate()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, utils.NewResponse("Failed to export users", nil))
	// 	return
	// }

	// c.Header("Content-Disposition", "attachment; filename=export.xlsx")
	// c.Data(http.StatusOK, "application/octet-stream", file.Bytes())
}

// @Summary Edit full name of user
// @Description Update only the full_name of a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body domain.UpdateFullNameRequest true "Update Full Name Request"
// @Success 200 {object} domain.SysUser
// @Router /users/update/{id} [put]
func (h *UserController) EditFullName(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewResponse("ok", "ok"))
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, utils.NewResponse("Invalid ID", nil))
	// 	return
	// }

	// var requestBody struct {
	// 	FullName string `json:"full_name"`
	// }

	// if err := c.ShouldBindJSON(&requestBody); err != nil {
	// 	c.JSON(http.StatusBadRequest, utils.NewResponse("Invalid request body", nil))
	// 	return
	// }

	// updatedUser, err := h.userUseCase.EditFullName(id, requestBody.FullName)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, utils.NewResponse("Failed to update full name", nil))
	// 	return
	// }

	// c.JSON(http.StatusOK, utils.NewResponse("Full name updated successfully", updatedUser))
}