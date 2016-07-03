package entities

// StatusResponse structure
type StatusResponse struct {
	Units    int64              `json:"units"`
	Start    string             `json:"start"`
	Stop     string             `json:"stop"`
	Duration int                `json:"duration"`
	Clients  map[string]*Client `json:"clients"`
	Parts    []Part             `json:"parts"`
}
