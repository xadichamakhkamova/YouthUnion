package handler

import (
	"api-gateway/internal/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
)

// @Router /events/ [post]
// @Summary Create a new event
// @Security BearerAuth
// @Tags Events
// @Accept json
// @Produce json
// @Param data body models.CreateEventRequest true "Event creation data"
// @Success 201 {object} models.Event
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) CreateEvent(c *gin.Context) {

	var req pb.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.CreateEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /events/{id} [put]
// @Summary Update an event
// @Security BearerAuth
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param data body models.UpdateEventRequest true "Event update data"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) UpdateEvent(c *gin.Context) {

	id := c.Param("id")
	var req pb.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.Id = id

	resp, err := h.service.UpdateEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /events/{id} [get]
// @Summary Get event by ID
// @Security BearerAuth
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
func (h *Handler) GetEvent(c *gin.Context) {

	id := c.Param("id")
	req := pb.GetEventRequest{Id: id}
	resp, err := h.service.GetEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /events/ [get]
// @Summary List events with filters
// @Security BearerAuth
// @Tags Events
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Param event_type query string false "Event type"
// @Param status query string false "Event status"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.ListEventsResponse
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) ListEvents(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.ListEventsRequest{
		Search:    c.Query("search"),
		EventType: pb.EventType(pb.EventType_value[c.DefaultQuery("event_type", "0")]),
		Status:    pb.EventStatus(pb.EventStatus_value[c.DefaultQuery("status", "0")]),
		Limit:     int32(limit),
		Page:      int32(page),
	}

	resp, err := h.service.ListEvents(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /events/{id} [delete]
// @Summary Delete event
// @Security BearerAuth
// @Tags Events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} models.DeleteEventResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) DeleteEvent(c *gin.Context) {

	id := c.Param("id")
	req := pb.DeleteEventRequest{Id: id}

	resp, err := h.service.DeleteEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
