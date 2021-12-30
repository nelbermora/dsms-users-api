package branch

import (
	"github.com/nelbermora/dsms-users-api/internal/model"
	"github.com/nelbermora/dsms-users-api/internal/person"
)

type Service interface {
	GetBranch(id int) (*model.Branch, error)
	GetUsersByBranch(int) ([]*model.User, error)
	GetBranches() ([]*model.Branch, error)
}

type service struct {
	br Repository
	ur person.UserRepository
}

func NewService(br Repository, ur person.UserRepository) Service {
	return &service{
		br: br,
		ur: ur,
	}
}

func (s *service) GetBranch(id int) (*model.Branch, error) {
	return s.br.GetBranch(id)
}

func (s *service) GetBranches() ([]*model.Branch, error) {
	return s.br.GetBranches(0)
}

func (s *service) GetUsersByBranch(branchId int) ([]*model.User, error) {
	return s.ur.GetUsersByBranch(branchId)
}
