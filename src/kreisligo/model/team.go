package model

import (

	"gorm.io/gorm"

)

type Team struct {
	gorm.Model
	LeagueID uint
	Name string `gorm:"notNull;size:50"`
	Roster []Player `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AwayGames []Game `gorm:"foreignKey:AwayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HomeGames []Game `gorm:"foreignKey:HomeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Wins uint `gorm:"-"` //COMPUTED 
	Losses uint `gorm:"-"` //COMPUTED
	Draws uint `gorm:"-"` //COMPUTED
	Points uint `gorm:"-"` //COMPUTED
}

// COMPUTATION FOR POINTS
func (t *Team) CalcPoints(tx *gorm.DB) (err error) {
	var wins uint;
	var losses uint;
	var draws uint;
	var points uint

	result := tx.Select("*").Where("team_id = ?", t.ID)
	if result.Error != nil {
		return result.Error
	}

	if len(t.HomeGames) < 1 && len(t.AwayGames) < 1 {
		return nil
	}
	for _, game := range t.HomeGames {
		if game.Status == "Beendet" || game.Status == "Läuft" {
			if game.Result == "Heim" {
				wins += 1
				points += 3
			} else if game.Result == "Auswärts" {
				losses += 1
			} else {
				draws += 1
				points += 1
			} 
		}
	}
	for _, game := range t.AwayGames {
		if game.Status == "Beendet" || game.Status == "Läuft" {
			if game.Result == "Heim" {
				losses += 1
			} else if game.Result == "Auswärts" {
				wins += 1
				points += 3
			} else {
				draws += 1
				points += 1
			}
		}
	}
	t.Wins = wins
	t.Losses = losses
	t.Draws = draws
	t.Points = points
	return nil
}

