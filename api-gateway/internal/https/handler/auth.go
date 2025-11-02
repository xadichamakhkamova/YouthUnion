package handler

import (
	"api-gateway/internal/https/token"
	"api-gateway/internal/models"
	"context"
	"net/http"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"

	"github.com/gin-gonic/gin"
)

// @Router /auth/register [post]
// @Summary Create User
// @Description Creates a new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User registration data"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) CreateUser(c *gin.Context) {

	var req pb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if req.Gender != "MALE" && req.Gender != "FEMALE" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid gender, must be MALE or FEMALE",
		})
		return
	}
	
	resp, err := h.service.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /auth/login [post]
// @Summary Get User by Identifier (Login)
// @Description Authenticates a user using identifier and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.GetUserByIdentifierRequest true "User credentials"
// @Success 200 {object} token.Tokens
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
func (h *Handler) GetUserByIdentifier(c *gin.Context) {

	var req pb.GetUserByIdentifierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.GetUserByIdentifier(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	token := token.GenereteJWTToken(resp)
	c.JSON(http.StatusOK, token)
}

// @Router /auth/change-password [patch]
// @Summary Change Password
// @Security BearerAuth
// @Description Allows user to change their password
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.ChangePasswordRequest true "Password change data"
// @Success 200 {object} models.ChangePasswordResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) ChangePassword(c *gin.Context) {

	var req pb.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.ChangePassword(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
