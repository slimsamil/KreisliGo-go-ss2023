package model

import (
	"gorm.io/gorm"
)

type Status string

const (
	ANSTEHEND Status = "Anstehend"
	LÄUFT Status = "Läuft"
	BEENDET Status = "Beendet"
)

type Game struct {
	gorm.Model

	LeagueID uint
	HomeID uint
	HomeGoals uint
	AwayID uint
	AwayGoals uint
	Events []Event `gorm:"foreignKey:GameID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Status Status `gorm:"notNull;type:ENUM('Anstehend', 'Läuft', 'Beendet')"`
	Result string //COMPUTED
}

// COMPUTATION FOR RESULTS
func (g *Game) FindResult(tx *gorm.DB) (err error) {
	result := tx.Select("*").Where("game_id = ?", g.ID)
	if result.Error != nil {
		return result.Error
	}
	if g.Status == BEENDET {
		var homeGoals uint = g.HomeGoals
		var awayGoals uint = g.AwayGoals
		var gameResult string;
		if g.Status == BEENDET {
			if homeGoals > awayGoals {
				gameResult = "Heim"
			} else if homeGoals < awayGoals {
				gameResult = "Auswärts"
			} else {
				gameResult = "Unentschieden"
			}
		}
		g.Result = gameResult
	}
	return nil
}

// COMPUTATION FOR EVENTS
func (g *Game) ComputateEvents(tx *gorm.DB) (err error) {
	var status Status;

	var homeGoals uint;
	var awayGoals uint;

	result := tx.Select("*").Where("game_id = ?", g.ID)
	if result.Error != nil {
		return result.Error
	}
	if len(g.Events) == 0 {
		return nil
	}
	if g.Events[0].EventType == "Anpfiff" {
		if g.Events[len(g.Events)-1].EventType == "Abpfiff" {
			status = "Beendet"
		} else {
			status = "Läuft"
		}
	} else {
		status = "Anstehend"
	}

	for _, event := range g.Events {
		if event.EventType == "Tor Heim" {
			homeGoals += 1;
		}
		if event.EventType == "Tor Auswärts" {
			awayGoals += 1;
		}
	}
	g.HomeGoals = homeGoals;
	g.AwayGoals = awayGoals;
	g.Status = status
	return nil
}

func (g *Game) ComputateGame(tx *gorm.DB) {
	g.ComputateEvents(tx)
	g.FindResult(tx)
}