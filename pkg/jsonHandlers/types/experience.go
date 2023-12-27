package types

type Experience struct {
	Experience []Job `json:"experience"`
}

type Job struct {
	Company      string   `json:"company"`
	Position     string   `json:"position"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Technologies []string `json:"technologies"`
	Description  string   `json:"description"`
}
