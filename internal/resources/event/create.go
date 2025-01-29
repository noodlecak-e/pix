package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noodlecak-e/pix/internal/models"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (e *Handler) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newEvent := models.Event{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	e.db.Create(&newEvent)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created event!",
		"event":   newEvent,
	})
}
