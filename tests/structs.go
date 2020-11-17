package tests

type Response struct {
	Message  string    `json:"message"`
	Success  bool      `json:"success"`
	Services []Service `json:"services"`
}

type Service struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Number     int    `json:"number"`
}
