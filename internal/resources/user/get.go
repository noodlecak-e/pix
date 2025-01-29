package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/models"
)

func (e *Handler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User
	tx := e.db.First(&user, "id = ?", id)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
