package entities

import "github.com/OuterInside/party/server/entities"

// StatusResponse structure
type StatusResponse struct {
	Units    int64              `json:"units"`
	Start    string             `json:"start"`
	Stop     string             `json:"stop"`
	Duration int                `json:"duration"`
	Clients  map[string]*Client `json:"clients"`
	Parts    []entities.Part    `json:"parts"`
}
