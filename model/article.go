package model

type Article struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	AuthorID int64  `json:"author_id"`
	Author   Author `json:"author"`
}

type InputNewArticle struct {
	Text     string `json:"text"`
	AuthorID int64  `json:"author_id"`
}

type InputOldArticle struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	AuthorID int64  `json:"author_id"`
}
