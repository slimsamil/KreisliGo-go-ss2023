package model

import ( 
	
	"gorm.io/gorm"

)

type Association struct {
	gorm.Model
	
	Name string `gorm:"notNull;size:50"`
	Leagues []League `gorm:"foreignKey:AssociationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}