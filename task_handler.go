package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	reqBody := r.Body
	if err := json.NewDecoder(reqBody).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Ошибка декодирования Task в TaskHandler %v \n", err)
		return
	}

	if task.Title == "" {
		http.Error(w, "У таски пустой title", http.StatusBadRequest)
		return
	}

	tx := t.db.Create(&task)
	if tx.Error != nil {
		http.Error(w, "DB Operation Error", http.StatusInternalServerError)
		log.Printf("Ошибка добавления Task в бд %v \n", tx.Error)
		return
	}

	// возвращаем response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := task

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Printf("Ошибка кодирования ответа в TaskHandler %v", err)
	}

	log.Printf("Задача успешно добавлена: ID=%d, Title=%s", task.ID, task.Title)
}

func (t *TaskHandler) GetTaskList(w http.ResponseWriter, _ *http.Request) {
	var taskList []Task
	if tx := t.db.Find(&taskList); tx.Error != nil {
		http.Error(w, "Ошибка поиска всех тасок в БД "+tx.Error.Error(), http.StatusInternalServerError)
		log.Printf("Ошибка поиска всех тасок в БД %v \n", tx.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&taskList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Ошибка енкодинга в TaskListHandler %v \n", err)
		return
	}

	log.Println("Список тасок успешно отдан клиенту")
}
