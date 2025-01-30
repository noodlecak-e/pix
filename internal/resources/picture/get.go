package picture

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
)

func (e *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("picture_id")

	var picture models.Picture
	tx := e.db.First(&picture, "id = ?", id)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"picture": picture,
	})
}
