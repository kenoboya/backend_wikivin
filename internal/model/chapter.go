package model

type Chapter struct {
	ID        int        `json:"chapter_id" db:"chapter_id"`
	ArticleID *int       `json:"article_id,omitempty" db:"article_id"`
	Name      string     `json:"name" db:"name"`
	Content   string     `json:"content" db:"content"`
	ParentID  *int       `json:"parent_id,omitempty" db:"parent_id"`
	Child     []*Chapter `json:"child,omitempty" db:"-"`
}
