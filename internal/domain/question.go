package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID       uint           `gorm:"primaryKey;autoIncrement"`
	Title    string         `gorm:"not null"`
	Type     string         `gorm:"not null"`
	Variants datatypes.JSON `gorm:"type:jsonb"`
	Answer   datatypes.JSON `gorm:"type:jsonb"`
	Tests    []Test         `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
