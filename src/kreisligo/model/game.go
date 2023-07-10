package model

import (
	"time"
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
	Events []Event `gorm:"foreignKey:GameID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Date time.Time `gorm:"notNull"`
	Status Status `gorm:"notNull;type:ENUM('Anstehend', 'Läuft', 'Beendet')"`
	Result string //COMPUTED
}

// COMPUTATION FOR RESULTS
func (g *Game) FindResult(tx *gorm.DB) (err error) {
	if g.Status == BEENDET {
		var gameResult string;
		result := tx.Select("*").Where("game_id = ?", g.ID)
		if result.Error != nil {
			return result.Error
		}
		if g.Status == BEENDET {
			if g.HomeGoals > g.AwayGoals {
				gameResult = "Heim"
			} else if g.HomeGoals < g.AwayGoals {
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
	var homeGoals uint;
	var awayGoals uint;
	var status Status;
	result := tx.Select("*").Where("game_id = ?", g.ID)
	if result.Error != nil {
		return result.Error
	}
	if g.Events[0].EventType == "Anpfiff" {
		if g.Events[-len(g.Events)-1].EventType == "Abpfiff" {
			status = "Beendet"
		} else {
			status = "Läuft"
		}
	} else {
		status = "Anstehend"
	}

	if status == "Läuft" {
		for _, event := range g.Events {
			if event.EventType == "Tor Heim" {
				homeGoals += 1;
			}
			if event.EventType == "Tor Auswärts" {
				awayGoals += 1;
			}
		}
	}

	g.AwayGoals = awayGoals
	g.HomeGoals = homeGoals
	g.Status = status
	return nil
}