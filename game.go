package model

type Event string

const (
	TOR Division = "TOR"
	GELB Division = "Gelbe Karte"
	ROT Division = "Rote Karte"
	WECHSEL Division = "AUSWECHSLUNG"
)

type Game struct {
	Home Team
	Away Team
	Events []Event
}