package person

import (
	"github.com/nelbermora/dsms-users-api/internal/model"
)

var CacheData map[int]*model.User = make(map[int]*model.User)

type UserRepository interface {
	GetUser(id int) (*model.User, error)
	UpdateUser(model.User) (*model.User, error)
	DeleteUser(int) error
	CreateUser(model.User) (*model.User, error)
	GetUsers() ([]model.User, error)
	GetUsersByBranch(branchId int) ([]*model.User, error)
}

type userRepository struct{}

func NewRepository() UserRepository {
	CacheData[1] = &model.User{
		ID:           1,
		Position:     "CEO",
		Name:         "John Doe",
		Department:   "T&O",
		StreetNumber: 123,
		Street:       "LongChamps",
		ZipCode:      "CP1022",
		City:         "Cork",
		Phone:        "041-123123",
		MobilePhone:  "+123-123-123",
		Email:        "john@doe.com",
		BranchId:     1,
		Login:        "2Factor",
		Status:       "Active",
	}
	return &userRepository{}
}

func (r *userRepository) GetUser(id int) (*model.User, error) {
	if user, ok := CacheData[id]; ok {
		return user, nil
	}
	return nil, nil
}

func (r *userRepository) UpdateUser(user model.User) (*model.User, error) {
	CacheData[user.ID] = &user
	return CacheData[user.ID], nil
}

func (r *userRepository) DeleteUser(id int) error {
	delete(CacheData, id)
	return nil
}

func (r *userRepository) CreateUser(user model.User) (*model.User, error) {
	user.ID = len(CacheData) + 1
	CacheData[user.ID] = &user
	return &user, nil
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	var data []model.User
	for key := range CacheData {
		data = append(data, *CacheData[key])
	}
	return data, nil
}

func (r *userRepository) GetUsersByBranch(branchId int) ([]*model.User, error) {
	var data []*model.User
	for key := range CacheData {
		data = append(data, CacheData[key])
	}
	return data, nil
}
