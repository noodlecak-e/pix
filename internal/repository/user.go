package repository

import "github.com/noodlecak-e/pix/internal/models"

func (e *Repository) UserExists(userID string) (bool, error) {
	var count int64
	result := e.db.Model(&models.User{}).Where("id = ?", userID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (e *Repository) CreateUser(user models.User) (models.User, error) {
	result := e.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (e *Repository) GetUsersPaginated(pagination Pagination) ([]models.User, error) {
	var users []models.User
	result := e.db.Offset(pagination.Offset()).Limit(pagination.Limit).Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}

func (e *Repository) GetUserByID(userID string) (models.User, error) {
	var user models.User
	result := e.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
