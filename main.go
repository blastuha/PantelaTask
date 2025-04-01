package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func TaskListHandler(w http.ResponseWriter, r *http.Request) {
	var taskList []Task
	if tx := DB.Find(&taskList); tx.Error != nil {
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

func TaskHandler(w http.ResponseWriter, r *http.Request) {
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

	tx := DB.Create(&task)
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

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<html><body><h1>Hello, World!</h1></body></html>`
	_, err := w.Write([]byte(html))
	if err != nil {
		fmt.Println("Ошибка записи ответа в HTMLHandler")
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/task-list", TaskListHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	router.HandleFunc("/api/html", HTMLHandler).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}
}
