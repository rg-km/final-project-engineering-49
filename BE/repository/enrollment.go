package repository

import (
	"belajar-golang/model/enrollment"
)

func (h *Repository) CreateEnrollmentAccessForLesson(UserID, LessonID int) error {

}

func (h *Repository) CheckEnrollmentKey(enrollmentRequest enrollment.EnrollmentRequest) (bool,error) {

}

func (h *Repository) CountEnrollmentEachUser(UserID int) (int, error) {
	
}

