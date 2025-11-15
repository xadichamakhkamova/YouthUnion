package handler

import (
	"api-gateway/internal/models"
	"context"
	"net/http"
	"strconv"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"

	"github.com/gin-gonic/gin"
)

// @Router /users/{id} [get]
// @Summary Get User By ID
// @Security BearerAuth
// @Description Returns user details by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) GetUserById(c *gin.Context) {

	id := c.Param("id")
	req := pb.GetUserByIdRequest{Id: id}

	resp, err := h.service.GetUserById(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /users/{id} [put]
// @Summary Update User
// @Security BearerAuth
// @Description Updates user information
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUserRequest true "User update data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var req pb.UpdateUserRequest
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /users [get]
// @Summary List Users
// @Security BearerAuth
// @Description Returns list of users
// @Tags Users
// @Produce json
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.UserList
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) ListUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.ListUsersRequest{Limit: int32(limit), Page: int32(page)}

	resp, err := h.service.ListUsers(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /users/{id} [delete]
// @Summary Delete User
// @Security BearerAuth
// @Description Deletes user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} models.DeleteUserResponse
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	req := pb.DeleteUserRequest{Id: id}

	resp, err := h.service.DeleteUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ! ------------------- User Roles -------------------
// @Router /users/{id}/roles [post]
// @Summary Assign Role to User
// @Security BearerAuth
// @Description Assigns a specific role to a user
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param role body models.AssignRoleRequest true "Role assignment data"
// @Success 201 {object} models.UserRole
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) AssignRoleToUser(c *gin.Context) {

	var req pb.AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.AssignRoleToUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /users/{id}/roles/{role_id} [delete]
// @Summary Remove Role from User
// @Security BearerAuth
// @Description Removes assigned role from user
// @Tags Roles
// @Param id path string true "User ID"
// @Param role_id path string true "Role ID"
// @Success 200 {object} models.RemoveRoleResponse
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) RemoveRoleFromUser(c *gin.Context) {

	id := c.Param("id")
	req := pb.RemoveRoleRequest{
		Id: id,
	}

	resp, err := h.service.RemoveRoleFromUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /users/{id}/roles [get]
// @Summary List User Roles
// @Security BearerAuth
// @Description Returns list of roles assigned to a user
// @Tags Roles
// @Param id path string true "User ID"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.UserRoleList
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) ListUserRoles(c *gin.Context) {

	userID := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.ListUserRolesRequest{UserId: userID, Page: int32(page), Limit: int32(limit)}

	resp, err := h.service.ListUserRoles(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ! ------------------- Role Handler -------------------
// @Router /roles [post]
// @Summary Create Role Type
// @Security BearerAuth
// @Description Creates a new role type (admin only)
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.CreateRoleRequest true "Role data"
// @Success 201 {object} models.RoleType
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) CreateRole(c *gin.Context) {

	var req pb.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := h.service.CreateRole(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Router /roles/{id} [get]
// @Summary Get Role By ID
// @Security BearerAuth
// @Description Returns role details by ID
// @Tags Roles
// @Param id path string true "Role ID"
// @Success 200 {object} models.RoleType
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) GetRoleById(c *gin.Context) {

	id := c.Param("id")
	req := pb.GetRoleByIdRequest{Id: id}

	resp, err := h.service.GetRoleById(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /roles/{id} [put]
// @Summary Update Role
// @Security BearerAuth
// @Description Updates role information
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param role body models.UpdateRoleRequest true "Role data"
// @Success 200 {object} models.RoleType
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) UpdateRole(c *gin.Context) {

	id := c.Param("id")
	var req pb.UpdateRoleRequest
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.UpdateRole(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /roles [get]
// @Summary List Roles
// @Security BearerAuth
// @Description Returns list of all role types
// @Tags Roles
// @Produce json
// @Success 200 {object} models.RoleTypeList
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) ListRoles(c *gin.Context) {

	req := pb.ListRolesRequest{}
	resp, err := h.service.ListRoles(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /roles/{id} [delete]
// @Summary Delete Role
// @Security BearerAuth
// @Description Deletes role type by ID
// @Tags Roles
// @Param id path string true "Role ID"
// @Success 200 {object} models.DeleteRoleResponse
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) DeleteRole(c *gin.Context) {

	id := c.Param("id")
	req := pb.DeleteRoleRequest{Id: id}

	resp, err := h.service.DeleteRole(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
