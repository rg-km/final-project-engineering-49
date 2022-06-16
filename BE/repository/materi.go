package repository

import (
	"belajar-golang/model/materi"
	"time"
)



func (h *Repository) FindAllMateri() ([]materi.Materi, int ,error) {
	var sqlStatement string
	var materis []materi.Materi
	sqlStatement = `SELECT id,title,contain,file_name,creator 
	FROM materis`
	rows, err:= h.db.Query(sqlStatement)
	if err != nil {
		return nil,0, err
	}
	var count int
	for rows.Next() {
		var materiTemp materi.Materi
		err := rows.Scan(&materiTemp.ID,&materiTemp.Title,&materiTemp.Contain,&materiTemp.FileName,&materiTemp.Creator)
		if err != nil {
			return nil,0,err
		}
		newMateri := materi.Materi{
			ID : materiTemp.ID,
			Title: materiTemp.Title,
			Contain: materiTemp.Contain,
			FileName: materiTemp.FileName,
			Creator: materiTemp.Creator,
		}
		materis = append(materis, newMateri)
		count++
	}
	rows.Close()
	return materis,count,nil
}


func (h *Repository) FindMateriByPage(Page int) ([]materi.Materi, int, error) {

	limit := 3
	offset := (Page - 1) * limit

	var sqlStatement string
	var materis []materi.Materi
	sqlStatement = `SELECT id,title,contain,file_name,creator 
	FROM materis LIMIT ? OFFSET ?`
	rows, err:= h.db.Query(sqlStatement,limit,offset)
	if err != nil {
		return nil,0, err
	}
	var count int
	for rows.Next() {
		var materiTemp materi.Materi
		err := rows.Scan(&materiTemp.ID,&materiTemp.Title,&materiTemp.Contain,&materiTemp.FileName,&materiTemp.Creator)
		if err != nil {
			return nil,0,err
		}
		newMateri := materi.Materi{
			ID : materiTemp.ID,
			Title: materiTemp.Title,
			Contain: materiTemp.Contain,
			FileName: materiTemp.FileName,
			Creator: materiTemp.Creator,
		}
		materis = append(materis, newMateri)
		count++
	}
	rows.Close()
	return materis,count,nil
}

func (h *Repository) FindMateriByFilter(materiFilterRequest materi.MateriFilterRequest) ([]materi.Materi, int, error) {
	
	limit := 3
	offset := (materiFilterRequest.Page - 1) * limit
	
	var sqlStatement string
	var materis []materi.Materi
	sqlStatement = `SELECT id,title,contain,file_name,creator 
	FROM materis WHERE title LIKE ? LIMIT ? OFFSET ?`
	rows, err:= h.db.Query(sqlStatement,"%"+materiFilterRequest.Keyword+"%",limit,offset)
	if err != nil {
		return nil,0, err
	}
	var count int
	for rows.Next() {
		var materiTemp materi.Materi
		err := rows.Scan(&materiTemp.ID,&materiTemp.Title,&materiTemp.Contain,&materiTemp.FileName,&materiTemp.Creator)
		if err != nil {
			return nil,0,err
		}
		newMateri := materi.Materi{
			ID : materiTemp.ID,
			Title: materiTemp.Title,
			Contain: materiTemp.Contain,
			FileName: materiTemp.FileName,
			Creator: materiTemp.Creator,
		}
		materis = append(materis, newMateri)
		count++
	}
	rows.Close()
	return materis,count,nil
}

//row := r.db.QueryRow(sqlStmt, nim)

func (h *Repository) FindMateriByID(MateriID int) (materi.Materi, error) {
	// TO DO Safa Auliya
	var sqlStatement string
	var materiTemp materi.Materi

	sqlStatement = `SELECT id,title,contain,file_name,creator 
	FROM materis WHERE id = ?`

	row := h.db.QueryRow(sqlStatement, MateriID)
	err := row.Scan(&materiTemp.ID,&materiTemp.Title,&materiTemp.Contain,&materiTemp.FileName,&materiTemp.Creator)
	if err != nil {
		return materiTemp, err
	}
	return materiTemp,nil
}

func (h *Repository) CreateMateri(materi materi.Materi) (materi.Materi, error) {
	// TO DO Safa Auliya
	sqlStmt := `INSERT INTO materis (title, contain,file_name,creator,created_at, updated_at) 
	VALUES (?,?,?,?,?,?);`
	datetime := time.Now()  
	_, err := h.db.Exec(sqlStmt, materi.Title,materi.Contain,materi.FileName,materi.Creator,datetime,datetime);
	if err != nil {
		return materi, err
	}
 
	return materi,nil
}

func (h *Repository) UpdateMateri(materi materi.Materi) (materi.Materi, error) {
	// TO DO Safa Auliya
	datetime := time.Now()  
	sqlStmt := `
		UPDATE materis
		SET title = ?, contain = ?, file_name = ?, creator = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := h.db.Exec(sqlStmt,materi.Title,materi.Contain,materi.FileName,materi.Creator,datetime,materi.ID)
	if err != nil {
		return materi,err
	}

	return materi,nil
}

func (h *Repository) DeleteMateri(ID int) error {
	// TO DO Safa Auliya
	sqlStmt := `DELETE FROM materis WHERE id = ?`
	_, err := h.db.Exec(sqlStmt, ID)
	if err != nil {
		return err
	}
	return nil
}