package usersService

import (
	"task1/internal/tasksService"
	api "task1/internal/web/users"

	"gorm.io/gorm"
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
	id := uint32(u.ID)
	email := u.Email

	tasks := make([]api.Task, len(u.Tasks))

	for i, modelTask := range u.Tasks {
		var responseTask api.Task

		taskID := uint32(modelTask.ID)
		userID := uint32(modelTask.UserID)

		responseTask.Id = taskID
		responseTask.UserId = &userID
		responseTask.Title = modelTask.Title
		responseTask.IsDone = modelTask.IsDone

		tasks[i] = responseTask
	}

	return api.UserResponse{
		Id:    &id,
		Email: &email,
		Tasks: &tasks,
	}
}
