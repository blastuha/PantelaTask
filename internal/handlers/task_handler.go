package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"task1/internal/dto"
	"task1/internal/models"
	"task1/internal/service"
)

type TaskHandler struct {
	service service.TasksService
}

func NewTaskHandler(service service.TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (t *TaskHandler) CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := t.service.CreateTask(&task); err != nil {
		if strings.Contains(err.Error(), "task has no title") {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, &task)
}

func (t *TaskHandler) GetTaskList(c echo.Context) error {
	allTasks, err := t.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var updateInput dto.TaskUpdateInput
	if err := c.Bind(&updateInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	updatedTask, err := t.service.UpdateTask(&updateInput, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, &updatedTask)
}

func (t *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	//uuid, err := uuid.Parse(id)
	//
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid UUID"})
	//}

	if err := t.service.DeleteTask(id); err != nil {
		if strings.Contains(err.Error(), "failed to find the task") {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
		}
		if strings.Contains(err.Error(), "failed to delete the task") {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}
		// Дефолтная ошибка
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An unexpected error occurred"})
	}

	return c.NoContent(http.StatusNoContent)
}
