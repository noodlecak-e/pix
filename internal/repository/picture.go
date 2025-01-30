package repository

import (
	"github.com/noodlecak-e/pix/internal/models"
)

func (e *Repository) GetPictureByID(pictureID string) (models.Picture, error) {
	var picture models.Picture
	result := e.db.First(&picture, "id = ?", pictureID)
	if result.Error != nil {
		return models.Picture{}, result.Error
	}
	return picture, nil
}

func (e *Repository) GetPicturesPaginated(pagination Pagination) ([]models.Picture, error) {
	var pictures []models.Picture
	result := e.db.Offset(pagination.Offset()).Limit(pagination.Limit).Find(&pictures)
	if result.Error != nil {
		return []models.Picture{}, result.Error
	}
	return pictures, nil
}

func (e *Repository) CreatePicture(picture models.Picture) (models.Picture, error) {
	result := e.db.Create(picture)
	if result.Error != nil {
		return models.Picture{}, result.Error
	}
	return picture, nil
}
