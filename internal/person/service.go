package person

import (
	"github.com/nelbermora/dsms-users-api/internal/model"
)

type Service interface {
	GetUser(id int) (*model.User, error)
	UpdateUser(model.User) (*model.User, error)
	DeleteUser(int) error
	CreateUser(model.User) (*model.User, error)
	GetUsers() ([]model.User, error)
}

type service struct {
	UserRepo UserRepository
}

func NewService(usrRepo UserRepository) Service {
	return &service{
		UserRepo: usrRepo,
	}
}

func (s *service) GetUser(id int) (*model.User, error) {
	usr, err := s.UserRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
func (s *service) UpdateUser(user model.User) (*model.User, error) {
	usr, err := s.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (s *service) DeleteUser(id int) error {
	err := s.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) CreateUser(user model.User) (*model.User, error) {
	usr, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (s *service) GetUsers() ([]model.User, error) {
	usrs, err := s.UserRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return usrs, nil
}
