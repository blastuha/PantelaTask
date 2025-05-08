package usersService

import (
	"fmt"
	"task1/internal/web/users"
)

type UsersService interface {
	GetAllUsers() ([]User, error)
	CreateUser(in users.CreateUserRequest) (*User, error)
	UpdateUser(in users.UpdateUserRequest) (*User, error)
	DeleteUser(id string) error
}

type usersService struct {
	repo UsersRepo
}

func NewUsersService(repo UsersRepo) UsersService {
	return &usersService{repo: repo}
}

func (u *usersService) GetAllUsers() ([]User, error) {
	userList, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("usersService.GetAllUsers: %w", err)
	}

	return userList, nil
}

func (u *usersService) CreateUser(in users.CreateUserRequest) (*User, error) {
	if in.Password == nil || len(*in.Password) < 6 {
		return nil, fmt.Errorf("usersService.CreateUser: password is required and must be at least 6 characters")
	}

	userToCreate := User{
		Email:    string(in.Email),
		Password: *in.Password,
	}

	createdUser, err := u.repo.CreateUser(&userToCreate)
	if err != nil {
		return nil, fmt.Errorf("usersService.CreateUser: %w", err)
	}

	return createdUser, nil
}

func (u *usersService) UpdateUser(in users.UpdateUserRequest) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}
