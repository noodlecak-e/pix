package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
)

func (e *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("event_id")

	var event models.Event
	tx := e.db.First(&event, "id = ?", id)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
