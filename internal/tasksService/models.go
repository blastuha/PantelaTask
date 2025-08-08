package tasksService

import (
	api "task1/internal/web/tasks"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"foreignKey:UserID"`
}

func (t *Task) ToResponse() api.CreateTask201JSONResponse {
	return api.CreateTask201JSONResponse{
		Id:     uint32(t.ID),
		Title:  t.Title,
		IsDone: t.IsDone,
	}
}
