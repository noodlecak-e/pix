package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/repository"
	"github.com/noodlecak-e/pix/pkg"
)

func (h *Handler) BulkGet(ctx *gin.Context) {
	pagination, err := repository.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}

	events, err := h.repository.GetEventsPaginated(*pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, repository.PaginatedResponse(events, pagination))
}
