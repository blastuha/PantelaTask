package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"task1/internal/models"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

func (t *TaskHandler) CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if task.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task has no title"})
	}

	tx := t.db.Create(&task)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create a task"})
	}

	return c.JSON(http.StatusOK, &task)
}

func (t *TaskHandler) GetTaskList(c echo.Context) error {
	var taskList []models.Task
	if tx := t.db.Find(&taskList); tx.Error != nil {
		log.Printf("Не удалось получить список тасок %v \n", tx.Error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get task list"})
	}
	return c.JSON(http.StatusOK, &taskList)
}
func (t *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	// формируем newTask из request
	var newTask models.Task
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Task который хотим обновить
	var oldTask models.Task
	if tx := t.db.First(&oldTask, id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// Отправляем в базу newTask
	if tx := t.db.Model(&oldTask).Updates(&newTask); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// возвращаем обновленного пользователя в ответе json
	return c.JSON(http.StatusOK, &newTask)
}

func (t *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	// Task который хотим удалить
	var taskToDelete models.Task
	if tx := t.db.First(&taskToDelete, id); tx.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Удаление Task
	if err := t.db.Delete(&taskToDelete).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// уст. статус ответа и сообщение
	return c.NoContent(http.StatusNoContent)
}
