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

	TeamId uint
	Name string `gorm:"notNull"`
	Position Position `gorm:"notNull"`
	JerseyNumber uint8 `gorm:"size:99"`
	Goals uint //COMPUTED
	YellowCards uint //COMPUTED
	RedCards uint //COMPUTED
}