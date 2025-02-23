package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noodlecak-e/pix/internal/models"
	"github.com/noodlecak-e/pix/pkg"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (e *Handler) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse(err))
		return
	}

	eventID := ctx.Param("event_id")

	newUser := models.User{
		ID:      uuid.New().String(),
		Name:    sql.NullString{String: req.Name, Valid: true},
		EventID: sql.NullString{String: eventID, Valid: true},
	}

	newUser, err := e.repository.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created user!",
		"user":    newUser,
	})
}
