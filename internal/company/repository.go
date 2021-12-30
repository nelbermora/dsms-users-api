package company

import "github.com/nelbermora/dsms-users-api/internal/model"

type Repository interface {
	GetCompany() (*model.Company, error)
}

type repository struct{}

var CacheCompany map[int]*model.Company = make(map[int]*model.Company)

func NewCompanyRepo() Repository {
	CacheCompany[1] = &model.Company{
		ID:           1,
		Name:         "Testing Company",
		Street:       "Baker Street",
		StreetNumber: 221,
		Phone:        "+54911-123-123",
		Email:        "company@testing.eu",
		ZipCode:      "NW16XE",
		City:         "New Jersey",
		State:        "Demo",
		Country:      "Demo Republic",
	}
	return &repository{}
}

func (br *repository) GetCompany() (*model.Company, error) {
	if user, ok := CacheCompany[1]; ok {
		return user, nil
	}
	return nil, nil
}
