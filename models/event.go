package models

type Event struct {
	ID                int    `json:"id"`
	SuggestedEventID  *int   `json:"suggested_event_id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	ImageURL          string `json:"image_url"`
	System            string `json:"system"`
	HostName          string `json:"host_name"`
	Host              *int   `json:"host"`
	RoomName          string `json:"room_name"`
	PuljeName         string `json:"pulje_name"`
	MaxPlayers        int    `json:"max_players"`
	ChildFriendly     bool   `json:"child_friendly"`
	AdultsOnly        bool   `json:"adults_only"`
	BeginnerFriendly  bool   `json:"beginner_friendly"`
	ExperiencedOnly   bool   `json:"experienced_only"`
	CanBeRunInEnglish bool   `json:"can_be_run_in_english"`
	LongRunning       bool   `json:"long_running"`
	ShortRunning      bool   `json:"short_running"`
	InsertedTime      string `json:"inserted_time"`
}
