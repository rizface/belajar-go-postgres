package web

type ArticleRequest struct {
	User_Id int    `json:"user_id"`
	Title   string `validate:"required" json:"title"`
	Content string `validate:"required" json:"content"`
}
