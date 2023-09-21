package repository

import (
	"github.com/test/models"
)

func CreateEvent(event models.RequiredEventFields) (models.Event, error) {
	// проверки полей нужно добавить
	newEvent := models.Event{
		EventName:        event.EventName,
		EventTheme:       event.EventTheme,
		EventLocation:    event.EventLocation,
		EventData:        event.EventData,
		EventBanner:      nil,
		EventDesctiption: []string{},
		EventPublic:      event.EventPublic,
		EventMembers:     0,
		EventChat:        models.EventChat{},
		EventAdmins:      models.EventAdmins{},
	}

	return newEvent, nil
}
