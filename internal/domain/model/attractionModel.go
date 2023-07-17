package model

type Attraction struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	LandName  string   `json:"land_name"`
	History   string   `json:"history"`
	CreatedIn string   `json:"created_in"`
	Creator   string   `json:"creator"`
	FunFacts  []string `json:"fun_facts"`
}
