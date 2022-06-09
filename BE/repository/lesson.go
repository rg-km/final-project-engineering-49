package repository

import (
	"belajar-golang/model/lesson"
)

func (h *Repository) CountLesson() (int,error){
	return 0,nil
} 

func (h *Repository) FindLessonByUser(UserID int,Page int) ([]lesson.LessonByUser,int,error) {	
	return nil,0,nil
}

func (h *Repository) FindLessonByPage(Page int) ([]lesson.Lesson,int,error) {	
	return nil,0,nil	
}

func (h *Repository) FindLessonByFilter(filterLessonRequest lesson.FilterLessonRequest) ([]lesson.Lesson,int,error)  {
	return nil,0,nil
}

func (h *Repository) FindLessonByID(LessonID int)(lesson.Lesson,error){
	return lesson.Lesson{},nil
}

func (h *Repository) CreateLesson(lesson lesson.Lesson)(lesson.Lesson, error){
	return lesson,nil
}

func (h *Repository) UpdateLesson(lesson lesson.Lesson)(lesson.Lesson, error){
	return lesson,nil
}

func (h *Repository) DeleteLesson(LessonID int)(error){ 
	return nil
}

func (h *Repository) IsAccessToLesson(userID,LessonID int) (bool) {
	return false
}