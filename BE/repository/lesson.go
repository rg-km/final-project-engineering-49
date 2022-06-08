package repository

import (
	"belajar-golang/model/lesson"
)

func (h *Repository) CountLesson() (int,error){

} 

func (h *Repository) FindLessonByUser(UserID int,Page int) ([]lesson.LessonByUser,int,error) {	
	
}

func (h *Repository) FindLessonByPage(Page int) ([]lesson.Lesson,int,error) {	
	
}

func (h *Repository) FindLessonByFilter(filterLessonRequest lesson.FilterLessonRequest) ([]lesson.Lesson,int,error)  {
	
}

func (h *Repository) FindLessonByID(LessonID int)(lesson.Lesson,error){
	
}

func (h *Repository) CreateLesson(lesson lesson.Lesson)(lesson.Lesson, error){

}

func (h *Repository) UpdateLesson(lesson lesson.Lesson)(lesson.Lesson, error){
	
}

func (h *Repository) DeleteLesson(LessonID int)(error){ 
	
}

func (h *Repository) IsAccessToLesson(userID,LessonID int) (bool) {

}