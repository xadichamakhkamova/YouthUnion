package handler

import (
	"api-gateway/internal/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
)

// @Router /scoring/give-score [post]
// @Summary Give score to a user
// @Security BearerAuth
// @Tags Scoring
// @Accept json
// @Produce json
// @Param data body models.GiveScoreRequest true "Score data"
// @Success 201 {object} models.Score
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *Handler) GiveScore(c *gin.Context) {

	var req pb.GiveScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := h.service.GiveScore(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /scoring/event/{event_id} [get]
// @Summary Get scores by event ID
// @Security BearerAuth
// @Tags Scoring
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.ScoreList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetScoresByEvent(c *gin.Context) {

	eventID := c.Param("event_id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.GetScoresByEventRequest{EventId: eventID, Limit: int32(limit), Page: int32(page)}
	resp, err := h.service.GetScoresByEvent(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /scoring/user/{user_id} [get]
// @Summary Get scores by user ID
// @Security BearerAuth
// @Tags Scoring
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.ScoreList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetScoresByUser(c *gin.Context) {

	userID := c.Param("user_id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.GetScoresByUserRequest{UserId: userID, Limit: int32(limit), Page: int32(page)}
	resp, err := h.service.GetScoresByUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /scoring/team/{team_id} [get]
// @Summary Get scores by team ID
// @Security BearerAuth
// @Tags Scoring
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.ScoreList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetScoresByTeam(c *gin.Context) {

	teamID := c.Param("team_id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.GetScoresByTeamRequest{TeamId: teamID, Limit: int32(limit), Page: int32(page)}
	resp, err := h.service.GetScoresByTeam(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /scoring/ranking [get]
// @Summary Get global ranking list
// @Security BearerAuth
// @Tags Scoring
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} models.RankingList
// @Failure 400 {object} models.ErrorResponse
func (h *Handler) GetGlobalRanking(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req := pb.GetGlobalRankingRequest{Limit: int32(limit), Page: int32(page)}
	resp, err := h.service.GetGlobalRanking(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
