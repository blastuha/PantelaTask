package usersService

import (
	"gorm.io/gorm"
	"task1/internal/tasksService"
	api "task1/internal/web/users"
)

type User struct {
	gorm.Model
	Email    string              `json:"email" gorm:"uniqueIndex;not null"`
	Password string              `json:"password" gorm:"not null"`
	Tasks    []tasksService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

func toUserResponse(u *User) api.UserResponse {
	if u == nil {
		return api.UserResponse{}
	}
	id := int64(u.ID)
	email := u.Email
	return api.UserResponse{
		Id:    &id,
		Email: &email,
	}
}
