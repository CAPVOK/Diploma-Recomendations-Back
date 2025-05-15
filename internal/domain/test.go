package domain

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	Deadline    time.Time  `gorm:"not null"`
	Courses     []Course   `gorm:"many2many:course_tests;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Questions   []Question `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
