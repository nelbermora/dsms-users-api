package model

type Company struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Street       string `json:"street"`
	StreetNumber int    `json:"streetNumber"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
}
