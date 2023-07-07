package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreateAssociation(association *model.Association) error {
	result := db.DB.Create(association)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new association with ID %v in database.", association.ID)
	log.Tracef("Stored: %v", association)
	return nil
}

func GetAssociations() ([]model.Association, error) {
	var associations []model.Association
	result := db.DB.Preload("Donations").Find(&associations)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", associations)
	return associations, nil
}

func GetAssociation(id uint) (*model.Association, error) {
	association := new(model.Association)
	result := db.DB.Preload("Donations").First(association, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", association)
	return association, nil
}

func UpdateAssociation(id uint, association *model.Association) (*model.Association, error) {
	existingAssociation, err := GetAssociation(id)
	if existingAssociation == nil || err != nil {
		return existingAssociation, err
	}
	existingAssociation.Name = association.Name
	existingAssociation.Leagues = association.Leagues
	result := db.DB.Save(existingAssociation)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated association.")
	entry.Tracef("Updated: %v", existingAssociation)
	return existingAssociation, nil
}

func DeleteAssociation(id uint) (*model.Association, error) {
	association, err := GetAssociation(id)
	if association == nil || err != nil {
		return association, err
	}
	result := db.DB.Delete(association)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted association.")
	entry.Tracef("Deleted: %v", association)
	return association, nil
}
