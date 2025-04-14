package service

import (
	"fmt"
	"strings"
	"task1/internal/dto"
	"task1/internal/models"
	"task1/internal/repository"
)

type tasksService struct {
	repo repository.TasksRepo
}

type TasksService interface {
	CreateTask(t *models.Task) error
	GetAllTasks() ([]models.Task, error)
	UpdateTask(updateData *dto.TaskUpdateInput, id string) (*models.Task, error)
	DeleteTask(id string) error
}

func NewTasksService(r repository.TasksRepo) TasksService {
	return &tasksService{repo: r}
}

func (s *tasksService) CreateTask(t *models.Task) error {
	if strings.TrimSpace(t.Title) == "" {
		return fmt.Errorf("task has no title")
	}
	if err := s.repo.CreateTask(t); err != nil {
		return fmt.Errorf("CreateTask: failed to create the task %w", err)
	}
	return nil
}

func (s *tasksService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *tasksService) UpdateTask(updateData *dto.TaskUpdateInput, id string) (*models.Task, error) {
	return s.repo.UpdateTask(updateData, id)
}

func (s *tasksService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
