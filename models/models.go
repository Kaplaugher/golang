package models

type Opportunity struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Client      string   `json:"client"`
	Files       []string `json:"files"`
	Needs       string   `json:"needs"`
}
