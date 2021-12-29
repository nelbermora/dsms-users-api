package model

type Branch struct {
	ID      int
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
}
