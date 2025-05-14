package domain

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	Users       []User `gorm:"many2many:user_courses;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tests       []Test `gorm:"many2many:course_tests;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
