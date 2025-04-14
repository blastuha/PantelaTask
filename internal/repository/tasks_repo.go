package repository

import (
	"fmt"
	"gorm.io/gorm"
	"task1/internal/dto"
	"task1/internal/models"
)

type TasksRepo interface {
	CreateTask(t *models.Task) error
	GetAllTasks() ([]models.Task, error)
	UpdateTask(t *dto.TaskUpdateInput, id string) (*models.Task, error)
	DeleteTask(id string) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TasksRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(t *models.Task) error {
	if err := r.db.Create(t).Error; err != nil {
		return fmt.Errorf("CreateTask: failed to create task: %w", err)
	}
	return nil
}

func (r *taskRepo) GetAllTasks() ([]models.Task, error) {
	var taskList []models.Task
	if err := r.db.Find(&taskList).Error; err != nil {
		return taskList, fmt.Errorf("GetAllTasks: failed to get tasks: %w", err)
	}
	return taskList, nil
}

func (r *taskRepo) UpdateTask(updateData *dto.TaskUpdateInput, id string) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return &task, fmt.Errorf("UpdateTask: failed to find task for updating: %w", err)
	}

	if err := r.db.Model(&task).Updates(&updateData).Error; err != nil {
		return &task, fmt.Errorf("UpdateTask: failed to update the task: %w", err)
	}

	return &task, nil
}

func (r *taskRepo) DeleteTask(id string) error {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return fmt.Errorf("DeleteTask: failed to find the task: %w", err)
	}

	if err := r.db.Delete(&task).Error; err != nil {
		return fmt.Errorf("DeleteTask: failed to delete the task: %w", err)
	}

	return nil
}
