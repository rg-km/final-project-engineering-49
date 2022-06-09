package repository

import (
	"belajar-golang/model/log"
)

func (h *Repository) CountLogEachUser(UserID int) (int,error) {
	return 0,nil
}

func (h *Repository) CreateOrUpdateLogForUser(log log.Log) error {
	return nil
}

func (h *Repository) GetScore(Answers []string,TestID []int)(int,error){
	return 0,nil
}
