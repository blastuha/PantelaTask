package dto

type TaskUpdateInput struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
