package entities

// EntryResponse structure
type EntryResponse struct {
	ID       string `json:"id"`
	Units    int64  `json:"units"`
	Part     int    `json:"part"`
	Start    string `json:"start"`
	Stop     string `json:"stop"`
	Duration int    `json:"duration"`
}
