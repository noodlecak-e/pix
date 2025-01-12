package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *EventHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "created event!",
	})
}
