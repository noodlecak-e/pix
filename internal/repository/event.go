package repository

import "github.com/noodlecak-e/pix/internal/models"

func (e *Repository) EventExists(eventID string) (bool, error) {
	var count int64
	result := e.db.Model(&models.Event{}).Where("id = ?", eventID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (e *Repository) GetEventsPaginated(pagination Pagination) ([]models.Event, error) {
	var events []models.Event
	result := e.db.Offset(pagination.Offset()).Limit(pagination.Limit).Find(&events)
	if result.Error != nil {
		return []models.Event{}, result.Error
	}
	return events, nil
}

func (e *Repository) CreateEvent(event models.Event) (models.Event, error) {
	result := e.db.Create(&event)
	if result.Error != nil {
		return models.Event{}, result.Error
	}
	return event, nil
}

func (e *Repository) GetEventByID(eventID string) (models.Event, error) {
	var event models.Event
	result := e.db.First(&event, "id = ?", eventID)
	if result.Error != nil {
		return models.Event{}, result.Error
	}
	return event, nil
}
