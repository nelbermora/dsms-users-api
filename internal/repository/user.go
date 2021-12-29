package repository

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
}

type userRepository struct{}

func NewRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUser(id int) (*model.User, error) {
	return CacheData[id], nil
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
