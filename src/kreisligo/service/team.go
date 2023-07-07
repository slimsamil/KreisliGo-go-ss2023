package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreateTeam(team *model.Team) error {
	result := db.DB.Create(team)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new team with ID %v in database.", team.ID)
	log.Tracef("Stored: %v", team)
	return nil
}

func GetTeams() ([]model.Team, error) {
	var teams []model.Team
	result := db.DB.Preload("Donations").Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", teams)
	return teams, nil
}

func GetTeam(id uint) (*model.Team, error) {
	team := new(model.Team)
	result := db.DB.Preload("Donations").First(team, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", team)
	return team, nil
}

func UpdateTeam(id uint, team *model.Team) (*model.Team, error) {
	existingTeam, err := GetTeam(id)
	if existingTeam == nil || err != nil {
		return existingTeam, err
	}
	existingTeam.Name = team.Name
	existingTeam.Roster = team.Roster
	existingTeam.Games = team.Games
	result := db.DB.Save(existingTeam)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated team.")
	entry.Tracef("Updated: %v", existingTeam)
	return existingTeam, nil
}

func DeleteTeam(id uint) (*model.Team, error) {
	team, err := GetTeam(id)
	if team == nil || err != nil {
		return team, err
	}
	result := db.DB.Delete(team)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted team.")
	entry.Tracef("Deleted: %v", team)
	return team, nil
}