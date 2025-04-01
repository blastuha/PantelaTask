package main

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

type TaskResponse struct {
	Task Task `json:"task"`
}
