package picture

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noodlecak-e/pix/internal/models"
	"github.com/noodlecak-e/pix/pkg"
	"github.com/noodlecak-e/pix/pkg/files"
)

func (e *Handler) Create(ctx *gin.Context) {
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

	multipartFile, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}
	defer multipartFile.Close()

	file, err := files.ConvertToFile(multipartFile, "/tmp/"+uuid.New().String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	fullPath, err := e.fileStorage.Create(ctx, []os.File{*file})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	newPicture := models.Picture{
		ID:        uuid.New().String(),
		EventID:   eventID,
		UserID:    userID,
		ImagePath: fullPath,
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
