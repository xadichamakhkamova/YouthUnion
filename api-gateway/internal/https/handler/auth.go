package handler

import (
	"api-gateway/internal/https/token"
	"context"
	"net/http"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"

	"github.com/gin-gonic/gin"
)

// ! ------------------- Authorization -------------------
// @Router /api/v1/auth/register [post]
// @Summary Register User
// @Description Registers a new user in the system
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /api/v1/auth/login [post]
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.GetUserByIdentifier(context.Background(), &req)
	if err != nil  || resp.Status != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token := token.GenereteJWTToken(int(req.Identifier))
	c.JSON(http.StatusOK, token)
}

// @Router /api/v1/auth/change-password [patch]
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.ChangePassword(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
