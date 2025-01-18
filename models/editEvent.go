package models

type EditEvent struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	Editing          bool   `json:"editing"`
	GameMaster       string `json:"game_master"`
	System           string `json:"system"`
	ShortDescription string `json:"short_description"`
}
