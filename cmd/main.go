package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"task1/config"
	"task1/internal/handlers"
	"task1/internal/models"
)

func main() {
	db, initDbErr := config.InitDB()
	if initDbErr != nil {
		log.Fatalf("Ошибка инициализации базы данных %v", db)
	}

	migrateErr := db.AutoMigrate(&models.Task{})
	if migrateErr != nil {
		log.Fatalf("Ошибка AutoMigrate %v", migrateErr)
	}

	taskHandler := handlers.NewTaskHandler(db)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/tasks", taskHandler.GetTaskList)
	//todo: сделать валидацию по title
	e.POST("/api/tasks", taskHandler.CreateTask)
	//todo: сделать валидацию по title
	e.PATCH("/api/tasks/:id", taskHandler.UpdateTask)
	//todo: проверить как у пантелы delete метод реализован, как будто у нас разные реализации
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	err := http.ListenAndServe(":8080", e)

	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}

}
