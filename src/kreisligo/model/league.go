package model

import (

	"gorm.io/gorm"

)

type Division string 

const (
	KREISLIGA Division = "Kreisliga"
	BEZIRKSLIGA Division = "Bezirksliga"
	LANDESLIGA Division = "Landesliga"
)

type League struct {
	gorm.Model

	AssociationId uint
	Name string `gorm:"notNull"`
	Division Division `gorm:"notNull"`
	Teams []Team `gorm:"foreignKey:LeagueID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// COMPUTATION FOR TABLE