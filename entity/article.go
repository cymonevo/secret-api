package entity

type Article struct {
	Title       string
	Description string
	Content     string
	Timestamp
}

type GetArticleResponse struct {
	Data Article
	Response
}
