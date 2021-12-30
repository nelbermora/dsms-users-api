package branch

import "github.com/nelbermora/dsms-users-api/internal/model"

type Repository interface {
	GetBranch(int) (*model.Branch, error)
	GetBranches(int) ([]*model.Branch, error)
}

type repository struct{}

func NewBranchRepo() Repository {
	CacheBranches[1] = &model.Branch{
		ID:           1,
		CompanyId:    1,
		ZipCode:      "NW16XE",
		City:         "New Jersey",
		Street:       "Baker Street",
		StreetNumber: 221,
		State:        "Demo",
		Country:      "Demo Republic",
		HeadQuarter:  true,
	}

	CacheBranches[2] = &model.Branch{
		ID:           2,
		CompanyId:    1,
		ZipCode:      "ASXQWE",
		City:         "London",
		Street:       "Sheen Street",
		StreetNumber: 774,
		State:        "Demo",
		Country:      "Demo Republic",
		HeadQuarter:  false,
	}

	CacheBranches[3] = &model.Branch{
		ID:           3,
		CompanyId:    1,
		ZipCode:      "CP1001",
		City:         "Mexico",
		Street:       "Liberacce",
		StreetNumber: 69,
		State:        "DF",
		Country:      "Mexico",
		HeadQuarter:  false,
	}

	return &repository{}
}

var CacheBranches map[int]*model.Branch = make(map[int]*model.Branch)

func (br *repository) GetBranch(id int) (*model.Branch, error) {
	if user, ok := CacheBranches[id]; ok {
		return user, nil
	}
	return nil, nil
}

func (br *repository) GetBranches(companyId int) ([]*model.Branch, error) {
	var data []*model.Branch
	for key := range CacheBranches {
		data = append(data, CacheBranches[key])
	}
	return data, nil
}
