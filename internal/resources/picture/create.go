package picture

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noodlecak-e/pix/internal/models"
)

type CreateRequest struct {
	ImageB64 string `json:"image_base64" binding:"required"`
}

func (e *Handler) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var exists bool
	if err := e.db.Model(&models.Event{}).Select("count(*) > 0").Where("id = ?", ctx.Param("event_id")).Find(&exists).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "event not found",
		})
		return
	}

	if err := e.db.Model(&models.User{}).Select("count(*) > 0").Where("id = ?", ctx.Param("user_id")).Find(&exists).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	newPicture := models.Picture{
		ID:          uuid.New().String(),
		EventID:     ctx.Param("event_id"),
		UserID:      ctx.Param("user_id"),
		ImageBase64: req.ImageB64,
	}

	tx := e.db.Create(&newPicture)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "saved picture!",
		"picture": newPicture,
	})
}
