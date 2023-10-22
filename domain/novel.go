package domain

type Novel struct {
	Id        string  `json:"id"`
	Code      string  `json:"code" binding:"required,len=4,alphanum"`
	Title     string  `json:"title" binding:"required,min=3,max=50"`
	Publisher string  `json:"publisher" binding:"required,min=3,max=50"`
	Year      string  `json:"year" binding:"required,numeric,len=4"`
	Author    string  `json:"author" binding:"required,min=3,max=100,alpha"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type Pagination struct {
	NovelList []Novel `json:"novel_list"`
	Page      int     `json:"page"`
	PerPage   int     `json:"per_page"`
	TotalPage int     `json:"total_page"`
	TotalData int     `json:"total_data"`
}
