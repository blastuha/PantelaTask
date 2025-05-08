package usersService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UsersRepo interface {
	GetAllUsers() ([]User, error)
	CreateUser(u *User) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id string) error
}

type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{db: db}
}

func (repo *usersRepo) GetAllUsers() ([]User, error) {
	var users []User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("usersRepo.GetAllUsers: %w", err)
	}

	return users, nil
}

func (repo *usersRepo) CreateUser(u *User) (*User, error) {
	if err := repo.db.Create(u).Error; err != nil {
		return nil, fmt.Errorf("usersRepo.CreateUser: %w", err)
	}

	return u, nil
}

func (repo *usersRepo) UpdateUser(u *User) (*User, error) {
	if err := repo.db.Save(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (repo *usersRepo) DeleteUser(id string) error {
	var user User

	if err := repo.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNoFound
		}

		return fmt.Errorf("userRepo.DeleteUser: %w", err)
	}

	if err := repo.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("userRepo.DeleteUser: %w", err)
	}

	return nil
}
