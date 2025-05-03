package tasksService

import (
	"gorm.io/gorm"
	"task1/internal/web/tasks"
)

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

func (t *Task) ToResponse() tasks.CreateTask201JSONResponse {
	return tasks.CreateTask201JSONResponse{
		Id:     int64(t.ID),
		Title:  t.Title,
		IsDone: t.IsDone,
	}
}
