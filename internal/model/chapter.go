package model

type Chapter struct {
	ID       int
	Name     string
	Content  string
	ParentID *int
	Child    []*Chapter
}