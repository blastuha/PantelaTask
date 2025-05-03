package tasksService

import (
	"fmt"
	"gorm.io/gorm"
)

type TasksRepo interface {
	CreateTask(t *Task) (*Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(t *tasksService.TaskUpdateInput, id string) (*Task, error)
	DeleteTask(id string) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TasksRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(t *Task) (*Task, error) {
	if err := r.db.Create(t).Error; err != nil {
		return nil, fmt.Errorf("CreateTask: failed to create task: %w", err)
	}
	return t, nil
}

func (r *taskRepo) GetAllTasks() ([]Task, error) {
	var taskList []Task
	if err := r.db.Find(&taskList).Error; err != nil {
		return taskList, fmt.Errorf("GetAllTasks: failed to get tasks: %w", err)
	}
	return taskList, nil
}

func (r *taskRepo) UpdateTask(updateData *tasksService.TaskUpdateInput, id string) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return &task, fmt.Errorf("UpdateTask: failed to find task for updating: %w", err)
	}

	task.Title = updateData.Title
	task.IsDone = updateData.IsDone

	if err := r.db.Save(&task).Error; err != nil {
		return &task, fmt.Errorf("UpdateTask: failed to save the task: %w", err)
	}

	return &task, nil
}

func (r *taskRepo) DeleteTask(id string) error {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return fmt.Errorf("DeleteTask: failed to find the task: %w", err)
	}

	if err := r.db.Delete(&task).Error; err != nil {
		return fmt.Errorf("DeleteTask: failed to delete the task: %w", err)
	}

	return nil
}
