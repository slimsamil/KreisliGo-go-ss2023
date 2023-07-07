package model

import (

	"gorm.io/gorm"

)

type Team struct {
	gorm.Model
	LeagueId uint
	Name string `gorm:"notNull;size:50"`
	Roster []Player `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Games []Game 
	Wins uint //COMPUTED
	Losses uint //COMPUTED
	Draws uint //COMPUTED
}

func (t *Team) CalcPoints(tx *gorm.DB) (err error) {
	this := *t
	for _, game := range t.Games {
		if game.Home.Name == this.Name{
			if game.Result == "Heim" {
				t.Wins += 1
			} else if game.Result == "Auswärts" {
				t.Losses += 1
			} else {
				t.Draws += 1
			}
		} else {
			if game.Result == "Heim" {
				t.Losses += 1
			} else if game.Result == "Auswärts" {
				t.Wins += 1
			} else {
				t.Draws += 1
			}
		}
	}
	return nil
}
