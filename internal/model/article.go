package model

type ArticlePage struct {
	Article  Article
	Chapters []Chapter
	InfoBox  InfoBox
}

type ArticleBriefInfo struct {
	ID          int
	Title       string
	LeadSection string
	Image       string
}

type Article struct {
	ArticleBriefInfo
	InfoBoxID  int
	ChaptersID []int
}