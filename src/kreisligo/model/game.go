package model

import "time"

type Event string

const (
	TOR Event = "Tor"
	GELB Event = "Gelbe Karte"
	GELBROT Event = "Gelb-rote Karte"
	ROT Event = "Rote Karte"
	WECHSEL Event = "Auswechslung"
)

type Game struct {
	Home Team
	Away Team
	Events []Event
	Date time.Time
}