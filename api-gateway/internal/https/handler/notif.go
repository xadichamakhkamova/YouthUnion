package handler

import (
	"api-gateway/internal/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
)

// @Router /notifications/send [post]
// @Summary Send a new notification
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Param data body models.SendNotificationRequest true "Notification send data"
// @Success 201 {object} models.SendNotificationResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) SendNotification(c *gin.Context) {

	var req pb.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.SendNotification(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
