package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"task1/config"
	"task1/internal/tasksService"
	"task1/internal/web/tasks"
)

func main() {
	db, initDbErr := config.InitDB()
	if initDbErr != nil {
		log.Fatalf("Ошибка инициализации базы данных %v", db)
	}

	tasksRepo := tasksService.NewTaskRepo(db)
	tasksService := tasksService.NewTasksService(tasksRepo)
	taskHandler := tasksService.NewTaskHandler(tasksService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	//e.GET("/api/tasks", taskHandler.GetTaskList)
	//e.POST("/api/tasks", taskHandler.CreateTask)
	//e.PATCH("/api/tasks/:id", taskHandler.UpdateTask)
	//e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	err := http.ListenAndServe(":8080", e)

	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}

}
