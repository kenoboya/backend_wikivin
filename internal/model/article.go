package model

type ArticlePage struct {
	Article  Article   `json:"Article"`
	Chapters []Chapter `json:"Chapters"`
	InfoBox  InfoBox   `json:"InfoBox"`
}

type Article struct {
	ID          int    `json:"article_id" db:"article_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
	AuthorID    int    `json:"author_id,omitempty" db:"author_id"`
}

type FavoriteArticle struct {
	AuthorID     int    `json:"author_id" db:"author_id"`
	AuthorImage  string `json:"author_image" db:"author_image"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	ArticleID    int    `json:"article_id" db:"article_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	ArticleImage string `json:"article_image" db:"article_image"`
}
