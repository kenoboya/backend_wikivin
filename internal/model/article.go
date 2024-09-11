package model

type ArticlePage struct {
	Article  Article   `json:"Article"`
	Chapters []Chapter `json:"Chapters"`
	InfoBox  InfoBox   `json:"InfoBox"`
}

type Article struct {
	ID          int    `json:"article_id" db:"article_id"`
	Title       string `json:"title" db:"title"`
	LeadSection string `json:"leadSection" db:"lead_section"`
	Image       string `json:"image" db:"image"`
}