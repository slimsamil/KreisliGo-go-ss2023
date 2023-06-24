package model

type Division string 

const (
	KREISLIGA Division = "Kreisliga"
	BEZIRKSLIGA Division = "Bezirksliga"
	LANDESLIGA Division = "Landesliga"
)

type League struct {

	Division Division
	Teams []Team
}