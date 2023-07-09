package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreateGame(game *model.Game) error {
	result := db.DB.Create(game)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new game with ID %v in database.", game.ID)
	log.Tracef("Stored: %v", game)
	return nil
}

func GetGames() ([]model.Game, error) {
	var games []model.Game
	result := db.DB.Preload("Donations").Find(&games)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", games)
	return games, nil
}

func GetGame(id uint) (*model.Game, error) {
	game := new(model.Game)
	result := db.DB.Preload("Donations").First(game, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", game)
	return game, nil
}

func UpdateGame(id uint, game *model.Game) (*model.Game, error) {
	existingGame, err := GetGame(id)
	if existingGame == nil || err != nil {
		return existingGame, err
	}
	existingGame.HomeID = game.HomeID
	existingGame.HomeGoals = game.HomeGoals
	existingGame.AwayID = game.AwayID
	existingGame.AwayGoals = game.AwayGoals
	existingGame.Events = game.Events
	existingGame.Status = game.Status
	existingGame.Date = game.Date
	result := db.DB.Save(existingGame)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated game.")
	entry.Tracef("Updated: %v", existingGame)
	return existingGame, nil
}

func DeleteGame(id uint) (*model.Game, error) {
	game, err := GetGame(id)
	if game == nil || err != nil {
		return game, err
	}
	result := db.DB.Delete(game)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted game.")
	entry.Tracef("Deleted: %v", game)
	return game, nil
}
