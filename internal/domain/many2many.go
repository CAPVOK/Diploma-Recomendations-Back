package domain

type UserCourse struct {
	UserID   uint `gorm:"primary_key"`
	CourseID uint `gorm:"primary_key"`
}

type CourseTest struct {
	CourseID uint `gorm:"primary_key"`
	TestID   uint `gorm:"primary_key"`
}

type TestQuestion struct {
	TestID     uint `gorm:"primary_key"`
	QuestionID uint `gorm:"primary_key"`
}
