package tasksService

import (
	"gorm.io/gorm"
	api "task1/internal/web/tasks"
)

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"foreignKey:UserID"`
}

func (t *Task) ToResponse() api.CreateTask201JSONResponse {
	return api.CreateTask201JSONResponse{
		Id:     int64(t.ID),
		Title:  t.Title,
		IsDone: t.IsDone,
	}
}
