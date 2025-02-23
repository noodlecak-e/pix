package picture

import (
	"github.com/noodlecak-e/pix/internal/repository"
	"github.com/noodlecak-e/pix/pkg/files"
)

type Handler struct {
	repository  repository.Repository
	fileStorage files.IFileIO
}

func NewHandler(repository repository.Repository, storage files.IFileIO) *Handler {
	return &Handler{
		repository:  repository,
		fileStorage: storage,
	}
}
