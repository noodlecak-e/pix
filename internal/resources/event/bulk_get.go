package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
	"github.com/noodlecak-e/pix/pkg"
)

func (h *Handler) BulkGet(ctx *gin.Context) {
	pagination, err := pkg.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}

	var events []models.Event
	result := h.db.Offset(pagination.Offset()).Limit(pagination.Limit).Find(&events)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(result.Error))
		return
	}

	ctx.JSON(http.StatusOK, pkg.PaginatedResponse(events, pagination))
}
