package handler

import (
	"api-gateway/internal/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
)

// @Router /teams/ [post]
// @Summary Create a new team
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param data body models.CreateTeamRequest true "Team creation data"
// @Success 201 {object} models.Team
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) CreateTeam(c *gin.Context) {

	var req pb.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.CreateTeam(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /teams/{id} [put]
// @Summary Update a team
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Param data body models.UpdateTeamRequest true "Team update data"
// @Success 200 {object} models.Team
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) UpdateTeam(c *gin.Context) {

	id := c.Param("id")
	var req pb.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.Id = id
	resp, err := h.service.UpdateTeam(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /teams/event/{event_id} [get]
// @Summary Get all teams for an event
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.TeamList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetTeamsByEvent(c *gin.Context) {

	eventID := c.Param("event_id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.GetTeamsByEventRequest{EventId: eventID, Limit: int32(limit), Page: int32(page)}

	resp, err := h.service.GetTeamsByEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /teams/{team_id}/members/{user_id} [delete]
// @Summary Remove a member from a team
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} models.StatusResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) RemoveTeamMember(c *gin.Context) {

	teamID := c.Param("team_id")
	userID := c.Param("user_id")

	req := pb.RemoveTeamMemberRequest{
		TeamId: teamID,
		UserId: userID,
	}

	resp, err := h.service.RemoveTeamMember(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /teams/{team_id}/members [get]
// @Summary Get all members of a team
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Success 200 {object} models.MemberList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetTeamMembers(c *gin.Context) {

	teamID := c.Param("team_id")
	req := pb.GetTeamRequest{TeamId: teamID}

	resp, err := h.service.GetTeamMembers(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /teams/invite [post]
// @Summary Invite a member to a team
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param data body models.InviteMemberRequest true "Invite a member"
// @Success 200 {object} models.InvitationsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) InviteMember(c *gin.Context) {

	var req pb.InviteMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.InviteMember(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /teams/respond [post]
// @Summary Respond to a team invitation
// @Security BearerAuth
// @Tags Teams
// @Accept json
// @Produce json
// @Param data body models.RespondInviteRequest true "Respond to invitation"
// @Success 200 {object} models.InvitationsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) RespondInvite(c *gin.Context) {

	var req pb.RespondInviteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.RespondInvite(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}