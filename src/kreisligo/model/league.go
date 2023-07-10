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

	AssociationID uint
	Name string `gorm:"notNull"`
	Division Division `gorm:"notNull;type:ENUM('Kreisliga', 'Bezirksliga', 'Landesliga')"`
	Teams []Team `gorm:"foreignKey:LeagueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Games []Game `gorm:"foreignKey:LeagueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// COMPUTATION FOR TABLE
func (l *League) SortTable(tx *gorm.DB) (err error) {
	result := tx.Model(&l.Teams).Select("*").Where("team_id = ?", l.ID).Order("Points")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
