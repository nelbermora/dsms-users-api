package company

import (
	"github.com/nelbermora/dsms-users-api/internal/branch"
	"github.com/nelbermora/dsms-users-api/internal/model"
)

type Service interface {
	GetCompany() (*model.Company, error)
	GetBranchesByCompany(int) ([]*model.Branch, error)
}

type service struct {
	cr Repository
	br branch.Repository
}

func NewService(cr Repository, br branch.Repository) Service {
	return &service{
		br: br,
		cr: cr,
	}
}

func (s *service) GetCompany() (*model.Company, error) {
	return s.cr.GetCompany()
}

func (s *service) GetBranchesByCompany(id int) ([]*model.Branch, error) {
	return s.br.GetBranches(id)
}
