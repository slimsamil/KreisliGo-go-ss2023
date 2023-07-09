package model

import (

	"gorm.io/gorm"

)

type Team struct {
	gorm.Model
	D uint
	Name string `gorm:"notNull;size:50"`
	Roster []Player `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AwayGames []Game `gorm:"foreignKey:AwayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HomeGames []Game `gorm:"foreignKey:HomeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Wins uint //COMPUTED
	Losses uint //COMPUTED
	Draws uint //COMPUTED
}

func (t *Team) CalcPoints(tx *gorm.DB) (err error) {
	for _, game := range t.HomeGames {
			if game.Result == "Heim" {
				t.Wins += 1
			} else if game.Result == "Auswärts" {
				t.Losses += 1
			} else {
				t.Draws += 1
			} 
		}

	for _, game := range t.AwayGames {
			if game.Result == "Heim" {
				t.Losses += 1
			} else if game.Result == "Auswärts" {
				t.Wins += 1
			} else {
				t.Draws += 1
			}
	}
	return nil
}

