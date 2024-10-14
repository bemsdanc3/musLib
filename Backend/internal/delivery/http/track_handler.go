package http

import (
	"Backend/internal/usecases"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TrackHandler struct {
	usecase usecases.TrackUsecase
}

func NewTrackHandler(usecase usecases.TrackUsecase) *TrackHandler {
	return &TrackHandler{usecase: usecase}
}

func (h *TrackHandler) GetTracksByUserID(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	tracks, err := h.usecase.GetTracksByUserID(userID)
	if err != nil {
		if err.Error() == "no tracks found for this user" {
			c.JSON(404, gin.H{"error": "No tracks found for this user"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(200, tracks)
}
