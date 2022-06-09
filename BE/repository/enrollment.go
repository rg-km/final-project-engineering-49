package repository

import (
	"belajar-golang/model/enrollment"
)

func (h *Repository) CreateEnrollmentAccessForLesson(UserID, LessonID int) error {

	return nil
}

func (h *Repository) CheckEnrollmentKey(enrollmentRequest enrollment.EnrollmentRequest) (bool,error) {
	return true,nil
}

func (h *Repository) CountEnrollmentEachUser(UserID int) (int, error) {
	return 0,nil
}

