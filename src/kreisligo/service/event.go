package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreateEvent(GameID uint, event *model.Event) error {
	event.GameID = GameID
	result := db.DB.Create(event)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new event with ID %v in database.", event.ID)
	log.Tracef("Stored: %v", event)
	return nil
}

func GetEvents() ([]model.Event, error) {
	var events []model.Event
	result := db.DB.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", events)
	return events, nil
}

func GetEvent(id uint) (*model.Event, error) {
	event := new(model.Event)
	result := db.DB.First(event, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", event)
	return event, nil
}

func UpdateEvent(id uint, event *model.Event) (*model.Event, error) {
	existingEvent, err := GetEvent(id)
	if existingEvent == nil || err != nil {
		return existingEvent, err
	}
	existingEvent.EventType = event.EventType
	existingEvent.PlayerID = event.PlayerID
	result := db.DB.Save(existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated event.")
	entry.Tracef("Updated: %v", existingEvent)
	return existingEvent, nil
}

func DeleteEvent(id uint) (*model.Event, error) {
	event, err := GetEvent(id)
	if event == nil || err != nil {
		return event, err
	}
	result := db.DB.Delete(event)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted event.")
	entry.Tracef("Deleted: %v", event)
	return event, nil
}
