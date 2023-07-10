package model

import(

	"gorm.io/gorm"

)

type EventType string

const (
	TOR EventType = "Tor"
	GELB EventType = "Gelbe Karte"
	GELBROT EventType = "Gelb-rote Karte"
	ROT EventType = "Rote Karte"
	WECHSEL EventType = "Auswechslung"
	SPIELBEGINN EventType = "Anpfiff"
	HALBZEIT EventType = "Halbzeit"
	ABPFIFF EventType = "Abpfiff"
)

type Event struct {
	gorm.Model

	GameID uint
	EventType EventType `gorm:"notNull;type:ENUM('Tor', 'Gelbe Karte', 'Gelb-rote Karte', 'Rote Karte', 'Auswechslung', 'Anpfiff', 'Halbzeit', 'Abpfiff')"`
	PlayerID uint
}