package materi

type MateriRequest struct {
	ID       int
	Title    string `json:"title" binding:"required"`
	Contain  string `json:"contain" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
	Creator  string `json:"creator" binding:"required"`
}

type MateriFilterRequest struct {
	Keyword string `json:"keyword" binding:"required"`
	Limit   int    `json:"limit" binding:"required"`
	Page    int    `json:"page" binding:"required"`
}
