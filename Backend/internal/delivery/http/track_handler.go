package http

import (
	"Backend/internal/entities"
	"Backend/internal/usecases"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TrackHandler struct {
	trackUsecase usecases.TrackUsecase
}

func NewTrackHandler(trackUsecase usecases.TrackUsecase) *TrackHandler {
	return &TrackHandler{trackUsecase: trackUsecase}
}

func (h *TrackHandler) GetTracksByUserID(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	tracks, err := h.trackUsecase.GetTracksByUserID(userID)
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

func (h *TrackHandler) GetAllTracks(c *gin.Context) {
	tracks, err := h.trackUsecase.GetAllTracks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tracks)
}

func (h *TrackHandler) CreateTrack(c *gin.Context) {
	var track entities.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.trackUsecase.CreateTrack(&track); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Track created successfully!"})
}
