package tasksService

import (
	"fmt"
	"strings"
	api "task1/internal/web/tasks"
)

type tasksService struct {
	repo TasksRepo
}

type TasksService interface {
	CreateTask(t *api.TaskCreateInput) (*Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(updateData *api.TaskUpdateInput, id string) (*Task, error)
	DeleteTask(id string) error
}

func NewTasksService(r TasksRepo) TasksService {
	return &tasksService{repo: r}
}

func (s *tasksService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *tasksService) CreateTask(t *api.TaskCreateInput) (*Task, error) {
	if strings.TrimSpace(t.Title) == "" {
		return nil, ErrInvalidInput
	}

	taskToCreate := Task{Title: t.Title, IsDone: t.IsDone, UserID: uint(t.UserId)}

	createdTask, err := s.repo.CreateTask(&taskToCreate)
	if err != nil {
		return nil, fmt.Errorf("CreateTask: failed to create the task: %w", err)
	}
	return createdTask, nil
}

func (s *tasksService) UpdateTask(updateData *api.TaskUpdateInput, id string) (*Task, error) {
	if strings.TrimSpace(updateData.Title) == "" {
		return nil, ErrInvalidInput
	}

	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = updateData.Title
	task.IsDone = updateData.IsDone

	return s.repo.UpdateTask(task)
}

func (s *tasksService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
