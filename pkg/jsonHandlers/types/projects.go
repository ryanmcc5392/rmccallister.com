package types

type Projects struct {
	ProjectName  string   `json:"projectName"`
	Technologies []string `json:"technologies"`
	Description  string   `json:"description"`
	Image        string   `json:"image"`
	URL          string   `json:"url"`
}
