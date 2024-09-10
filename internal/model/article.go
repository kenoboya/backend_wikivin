package model

type ArticlePage struct {
	Article  Article
	Chapters []Chapter
	InfoBox  InfoBox
}

type Article struct {
	ID          int
	Title       string
	LeadSection string
	Image       string
}