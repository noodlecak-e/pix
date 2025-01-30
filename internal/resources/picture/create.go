package picture

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noodlecak-e/pix/internal/models"
	"github.com/noodlecak-e/pix/pkg"
)

type CreateRequest struct {
	ImageB64 string `json:"image_base64" binding:"required"`
}

func (e *Handler) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}

	eventID := ctx.Param("event_id")
	userID := ctx.Param("user_id")

	eventExists, err := e.repository.EventExists(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}
	if !eventExists {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse(errors.New("event not found")))
		return
	}

	userExists, err := e.repository.UserExists(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}
	if !userExists {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse(errors.New("user not found")))
		return
	}

	newPicture := models.Picture{
		ID:          uuid.New().String(),
		EventID:     eventID,
		UserID:      userID,
		ImageBase64: req.ImageB64,
	}

	newPicture, err = e.repository.CreatePicture(newPicture)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "saved picture!",
		"picture": newPicture,
	})
}
