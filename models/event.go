package models

type Event struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	GameMaster       string `json:"game_master"`
	System           string `json:"system"`
}
