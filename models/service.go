package models

type Service struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Number     int    `json:"number"`
}
