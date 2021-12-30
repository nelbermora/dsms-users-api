package model

type Branch struct {
	ID           int    `json:"id"`
	CompanyId    int    `json:"companyId"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
	Street       string `json:"street"`
	StreetNumber int    `json:"streetNumber"`
	State        string `json:"state"`
	Country      string `json:"country"`
	HeadQuarter  bool   `json:"isHeadQuarter"`
}
