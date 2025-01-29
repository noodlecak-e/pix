package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
)

func (e *Handler) BulkGet(ctx *gin.Context) {
	var events []models.Event
	tx := e.db.Find(&events)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
