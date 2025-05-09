package usersService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task1/internal/web/users"
)

type UsersService interface {
	GetAllUsers() ([]User, error)
	CreateUser(in users.CreateUserRequest) (*User, error)
	UpdateUser(id string, in users.UpdateUserRequest) (*User, error)
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

func (u *usersService) UpdateUser(id string, in users.UpdateUserRequest) (*User, error) {
	var existingUser User
	if err := u.repo.(*usersRepo).db.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNoFound
		}
		return nil, fmt.Errorf("usersService.UpdateUser: %w", err)
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

	updatedUser, err := u.repo.UpdateUser(&existingUser)
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
