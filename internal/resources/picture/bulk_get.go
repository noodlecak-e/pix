package picture

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
	"github.com/noodlecak-e/pix/pkg"
)

func (e *Handler) BulkGet(ctx *gin.Context) {
	pagination, err := pkg.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}

	var pictures []models.Picture
	result := e.db.Offset(pagination.Offset()).Limit(pagination.Limit).Find(&pictures)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(result.Error))
		return
	}

	ctx.JSON(http.StatusOK, pkg.PaginatedResponse(pictures, pagination))
}
