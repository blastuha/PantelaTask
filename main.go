package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, initDbErr := InitDB()
	if initDbErr != nil {
		log.Fatalf("Ошибка инициализации базы данных %v", db)
	}

	migrateErr := db.AutoMigrate(&Task{})
	if migrateErr != nil {
		log.Fatalf("Ошибка AutoMigrate %v", migrateErr)
	}

	taskHandler := NewTaskHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/api/task-list", taskHandler.GetTaskList).Methods("GET")
	router.HandleFunc("/api/task", taskHandler.CreateTask).Methods("POST")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}

}
