package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/db"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
)

func CreatePlayer(TeamID uint, player *model.Player) error {
	player.TeamID = TeamID
	result := db.DB.Create(player)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new player with ID %v in database.", player.ID)
	log.Tracef("Stored: %v", player)
	return nil
}

func GetPlayers() ([]model.Player, error) {
	var players []model.Player
	result := db.DB.Preload("Events").Find(&players)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", players)
	return players, nil
}

func GetPlayer(id uint) (*model.Player, error) {
	player := new(model.Player)
	result := db.DB.Preload("Events").First(player, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", player)
	return player, nil
}

func UpdatePlayer(id uint, player *model.Player) (*model.Player, error) {
	existingPlayer, err := GetPlayer(id)
	if existingPlayer == nil || err != nil {
		return existingPlayer, err
	}
	existingPlayer.Name = player.Name
	existingPlayer.Position = player.Position
	existingPlayer.JerseyNumber = player.JerseyNumber
	result := db.DB.Save(existingPlayer)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated player.")
	entry.Tracef("Updated: %v", existingPlayer)
	return existingPlayer, nil
}

func DeletePlayer(id uint) (*model.Player, error) {
	player, err := GetPlayer(id)
	if player == nil || err != nil {
		return player, err
	}
	result := db.DB.Delete(player)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted player.")
	entry.Tracef("Deleted: %v", player)
	return player, nil
}
