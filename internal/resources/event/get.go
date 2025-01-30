package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/pkg"
)

func (e *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("event_id")

	event, err := e.repository.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
