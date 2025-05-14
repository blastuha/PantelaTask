package usersService

import (
	"errors"
	"fmt"
	"task1/internal/tasksService"
	"task1/internal/web/users"
)

type UsersService interface {
	GetAllUsers() ([]User, error)
	CreateUser(in users.CreateUserRequest) (*User, error)
	UpdateUser(id string, in users.UpdateUserRequest) (*User, error)
	DeleteUser(id string) error
	GetTasksForUser(id uint) ([]tasksService.Task, error)
}

type usersService struct {
	repo UsersRepo
}

func (u *usersService) GetTasksForUser(id uint) ([]tasksService.Task, error) {
	tasks, err := u.repo.GetTasksForUser(id)
	if errors.Is(err, ErrUserNoFound) {
		return nil, ErrUserNoFound
	}
	if err != nil {
		return nil, fmt.Errorf("usersService.GetTasksForUser: %w", err)
	}
	return tasks, nil
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

func (u *usersService) UpdateUser(id string, in users.UpdateUserRequest) (*User, error) {
	existingUser, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if in.Email != nil {
		existingUser.Email = string(*in.Email)
	}
	if in.Password != nil {
		if len(*in.Password) < 6 {
			return nil, fmt.Errorf("password must be at least 6 characters")
		}
		existingUser.Password = *in.Password
	}

	updatedUser, err := u.repo.UpdateUser(existingUser)
	if err != nil {
		return nil, fmt.Errorf("usersService.UpdateUser: %w", err)
	}

	return updatedUser, nil
}

func (u *usersService) DeleteUser(id string) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		if errors.Is(err, ErrUserNoFound) {
			return ErrUserNoFound
		}
		return fmt.Errorf("usersService.DeleteUser: %w", err)
	}

	return nil
}
