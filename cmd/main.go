package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"task1/config"
	"task1/internal/tasksService"
	"task1/internal/usersService"
	"task1/internal/web/tasks"
	"task1/internal/web/users"
)

func main() {
	db, initDbErr := config.InitDB()
	if initDbErr != nil {
		log.Fatalf("Ошибка инициализации базы данных %v", db)
	}

	tasksRepo := tasksService.NewTaskRepo(db)
	tasksSvc := tasksService.NewTasksService(tasksRepo)
	taskHandler := tasksService.NewTaskHandler(tasksSvc)

	usersRepo := usersService.NewUsersRepo(db)
	usersSvc := usersService.NewUsersService(usersRepo)
	usersHandler := usersService.NewUsersHandler(usersSvc)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	usersStrict := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrict)

	err := http.ListenAndServe(":8080", e)

	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}

}
