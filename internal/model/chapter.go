package model

type Chapter struct {
	ID        int
	ArticleID int
	Name      string
	Content   string
	ParentID  *int
	Child     []*Chapter
}