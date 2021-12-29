package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Title    string `json:"title"`
	BranchId int    `json:"branchId"`
	Role     int    `json:"role"`
	Status   string `json:"status"`
	Branch   Branch `json:"branch"`
}

func (u *User) Validate() error {
	return nil
}
