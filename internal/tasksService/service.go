package tasksService

import (
	"fmt"
	"strings"
)

type tasksService struct {
	repo TasksRepo
}

type TasksService interface {
	CreateTask(t *TaskCreateInput) (*Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(updateData *TaskUpdateInput, id string) (*Task, error)
	DeleteTask(id string) error
}

func NewTasksService(r TasksRepo) TasksService {
	return &tasksService{repo: r}
}

func (s *tasksService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *tasksService) CreateTask(t *TaskCreateInput) (*Task, error) {
	if strings.TrimSpace(t.Title) == "" {
		return nil, fmt.Errorf("task has no title")
	}

	taskToCreate := Task{Title: t.Title, IsDone: t.IsDone}

	createdTask, err := s.repo.CreateTask(&taskToCreate)
	if err != nil {
		return nil, fmt.Errorf("CreateTask: failed to create the task: %w", err)
	}
	return createdTask, nil
}

func (s *tasksService) UpdateTask(updateData *TaskUpdateInput, id string) (*Task, error) {
	return s.repo.UpdateTask(updateData, id)
}

func (s *tasksService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
