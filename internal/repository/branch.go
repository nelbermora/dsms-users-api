package repository

import "github.com/nelbermora/dsms-users-api/internal/model"

type BranchRepository interface {
	GetBranch(int) (model.Branch, error)
}

type branchRepository struct{}

func NewBranchRepo() BranchRepository {
	return &branchRepository{}
}

func (br *branchRepository) GetBranch(int) (model.Branch, error) {
	return model.Branch{
		Name:   "Testing Branch",
		Status: "Active",
		ID:     123,
	}, nil
}
