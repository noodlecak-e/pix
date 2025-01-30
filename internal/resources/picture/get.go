package picture

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/pkg"
)

func (e *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("picture_id")

	picture, err := e.repository.GetPictureByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"picture": picture,
	})
}
