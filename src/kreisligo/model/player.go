package model

import (

	"gorm.io/gorm"
	
)

type Position string 

const (
	TORWART Position = "Torwart"
	VERTEIDIGUNG Position = "Verteidigung"
	MITTELFELD Position = "Mittelfeld"
	STURM Position = "Sturm"
)

type Player struct {
	gorm.Model

	TeamID uint
	Name string `gorm:"notNull"`
	Position Position `gorm:"notNull;type:ENUM('Torwart', 'Verteidigung', 'Mittelfeld', 'Sturm')"`
	JerseyNumber uint8 `gorm:"size:99"`
	Events []Event `gorm:"foreignKey:PlayerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Goals uint //COMPUTED
	YellowCards uint //COMPUTED
	RedCards uint //COMPUTED
}

// COMPUTATION FOR EVENTS
func (p *Player) CalcStats(tx *gorm.DB) (err error) {
	var goals uint;
	var yellowCards uint;
	var redCards uint;
	result := tx.Select("*").Where("player_id = ?", p.ID)
	if result.Error != nil {
		return result.Error
	}
	if len(p.Events) == 0 {
		return nil
	}

	for _, event := range p.Events {
		if(event.EventType == "Tor Heim" || event.EventType == "Tor Auswärts"){
			goals += 1;
		}
		if(event.EventType == "Gelbe Karte"){
			yellowCards += 1;
		}
		if(event.EventType == "Rote Karte" || event.EventType == "Gelb-rote Karte"){
			redCards += 1;
		}
	}
	p.Goals = goals;
	p.YellowCards = yellowCards;
	p.RedCards = redCards;
	return nil;
}