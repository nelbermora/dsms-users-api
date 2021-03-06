package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Department   string `json:"department"`
	Street       string `json:"street"`
	StreetNumber int    `json:"streetNumber"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
	Phone        string `json:"phone"`
	MobilePhone  string `json:"mobilePhone"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	BranchId     int    `json:"branchId"`
	Login        string `json:"login"`
	Status       string `json:"status"`
}

func (u *User) Validate() error {
	return nil
}
