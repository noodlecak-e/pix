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

	tx := e.db.Create(&newEvent)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created event!",
		"event":   newEvent,
	})
}
