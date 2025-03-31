package main

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
