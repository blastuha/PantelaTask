package tasksService

type TaskUpdateInput struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

type TaskCreateInput struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
