package tasksService

import (
	"context"
	"errors"
	api "task1/internal/web/tasks"
)

type TaskHandler struct {
	service TasksService
}

func (t *TaskHandler) GetTaskList(_ context.Context, _ api.GetTaskListRequestObject) (api.GetTaskListResponseObject, error) {
	allTasks, err := t.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := api.GetTaskList200JSONResponse{}

	for _, tsk := range allTasks {
		task := api.Task{Id: int64(tsk.ID), Title: tsk.Title, IsDone: tsk.IsDone}
		response = append(response, task)
	}

	return response, nil
}

func (t *TaskHandler) CreateTask(_ context.Context, request api.CreateTaskRequestObject) (api.CreateTaskResponseObject, error) {
	taskRequest := request.Body
	var response api.CreateTask201JSONResponse

	taskToCreate := api.TaskCreateInput{Title: taskRequest.Title, IsDone: taskRequest.IsDone, UserId: taskRequest.UserId}

	createdTask, err := t.service.CreateTask(&taskToCreate)

	if err != nil {
		if errors.Is(err, ErrInvalidInput) {
			return api.CreateTask400JSONResponse{Error: "Task has no title"}, nil
		}

		return nil, err
	}

	response = createdTask.ToResponse()

	return response, nil
}

func (t *TaskHandler) DeleteTask(_ context.Context, request api.DeleteTaskRequestObject) (api.DeleteTaskResponseObject, error) {
	id := request.Id
	err := t.service.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return api.DeleteTask204Response{}, nil
}

func (t *TaskHandler) UpdateTask(_ context.Context, request api.UpdateTaskRequestObject) (api.UpdateTaskResponseObject, error) {
	requestBody := request.Body
	requestId := request.Id

	taskToUpdate := api.TaskUpdateInput{
		Title:  requestBody.Title,
		IsDone: requestBody.IsDone,
	}

	updatedTask, err := t.service.UpdateTask(&taskToUpdate, requestId)
	if err != nil {

		switch {
		case errors.Is(err, ErrTaskNoFound):
			return api.UpdateTask404JSONResponse{Error: "Task not found"}, nil
		case errors.Is(err, ErrInvalidInput):
			return api.UpdateTask400JSONResponse{Error: "Task has no title"}, nil
		default:
			return nil, err
		}

	}

	response := api.UpdateTask200JSONResponse{
		Id:     int64(updatedTask.ID),
		Title:  updatedTask.Title,
		IsDone: updatedTask.IsDone,
	}

	return response, nil
}

func NewTaskHandler(service TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}
