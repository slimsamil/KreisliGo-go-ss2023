package model

import (
	"gorm.io/gorm"
	"time"
)

type Status string

const (
	ANSTEHEND Status = "Anstehend"
	LÄUFT Status = "Läuft"
	BEENDET Status = "Beendet"
)

type Game struct {
	gorm.Model
	Home Team `gorm:"notNull"`
	HomeGoals uint
	Away Team `gorm:"notNull"`
	AwayGoals uint
	Events []Event `gorm:"foreignKey:GameID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Date time.Time `gorm:"notNull"`
	Status Status `gorm:"notNull"`
	Result string //COMPUTED
}

func (g *Game) FindResult(tx *gorm.DB) (err error) {

	if g.Status == BEENDET {
		if g.HomeGoals > g.AwayGoals {
			g.Result = "Heim"
		} else if g.HomeGoals < g.AwayGoals {
			g.Result = "Auswärts"
		} else {
			g.Result = "Unentschieden"
		}
	}
	return nil
}

// COMPUTATION FOR EVENTS
