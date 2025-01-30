package picture

import "github.com/noodlecak-e/pix/internal/repository"

type Handler struct {
	repository repository.Repository
}

func NewHandler(repository repository.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}
