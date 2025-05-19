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

type UserTests struct {
	TestID uint           `gorm:"primary_key"`
	UserID uint           `gorm:"primary_key"`
	Status UserTestStatus `gorm:"type:varchar(20);not null;status IN ('READY_FOR_PRORGESS', 'IN_PROGRESS', 'COMPLETED');default:'READY_FOR_PROGRESS'"`
}

type UserTestStatus string

const (
	ReadyForProgress UserTestStatus = "READY_FOR_PROGRESS"
	InProgress       UserTestStatus = "IN_PROGRESS"
	Completed        UserTestStatus = "COMPLETED"
)
