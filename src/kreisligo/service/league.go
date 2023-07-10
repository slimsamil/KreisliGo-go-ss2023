package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreateLeague(associationID uint, league *model.League) error {
	league.AssociationID = associationID
	result := db.DB.Create(league)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new league with ID %v in database.", league.ID)
	log.Tracef("Stored: %v", league)
	return nil
}

func GetLeagues() ([]model.League, error) {
	var leagues []model.League
	result := db.DB.Preload("Teams").Find(&leagues)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", leagues)
	return leagues, nil
}

func GetLeague(id uint) (*model.League, error) {
	league := new(model.League)
	result := db.DB.Preload("Teams").First(league, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", league)
	return league, nil
}

func UpdateLeague(id uint, league *model.League) (*model.League, error) {
	existingLeague, err := GetLeague(id)
	if existingLeague == nil || err != nil {
		return existingLeague, err
	}
	existingLeague.Name = league.Name
	existingLeague.Division = league.Division
	existingLeague.Teams = league.Teams
	result := db.DB.Save(existingLeague)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated league.")
	entry.Tracef("Updated: %v", existingLeague)
	return existingLeague, nil
}

func DeleteLeague(id uint) (*model.League, error) {
	league, err := GetLeague(id)
	if league == nil || err != nil {
		return league, err
	}
	result := db.DB.Delete(league)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted league.")
	entry.Tracef("Deleted: %v", league)
	return league, nil
}
