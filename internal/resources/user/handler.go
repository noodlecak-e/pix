package user

import (
	"github.com/noodlecak-e/pix/internal/repository"
)

type Handler struct {
	repository repository.Repository
}

func NewHandler(respository repository.Repository) *Handler {
	return &Handler{
		repository: respository,
	}
}
