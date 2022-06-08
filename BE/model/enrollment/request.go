package enrollment

type EnrollmentRequest struct {
	LessonID  int    `binding:"required"`
	EnrollKey string `binding:"required"`
}