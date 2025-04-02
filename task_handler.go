package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	// уст. статус
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

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// формируем newTask из request
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}

	// Task который хотим обновить
	var oldTask Task
	if tx := t.db.First(&oldTask, id); tx.Error != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}

	// Отправляем в базу newTask
	if tx := t.db.Model(&oldTask).Updates(&newTask); tx.Error != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}

	// возвращаем обновленного пользователя в ответе json
	if err := json.NewEncoder(w).Encode(&oldTask); err != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Task который хотим удалить
	var taskToDelete Task
	if tx := t.db.First(&taskToDelete, id); tx.Error != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}

	// Удаление Task
	if err := t.db.Delete(&taskToDelete).Error; err != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}

	// уст. статус ответа и сообщение
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Удаление прошло успешно"})
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusBadRequest)
		return
	}
}
