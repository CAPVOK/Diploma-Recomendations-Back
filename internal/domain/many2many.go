package domain

type UserCourse struct {
	UserID   uint
	CourseID uint
}

type CourseTest struct {
	CourseID uint
	TestID   uint
}

type TestQuestion struct {
	TestID     uint
	QuestionID uint
}
