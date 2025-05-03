package tasksService

import (
	"context"
	"task1/internal/web/tasks"
)

type TaskHandler struct {
	service tasksService.TasksService
}

func (t *TaskHandler) GetTaskList(_ context.Context, _ tasks.GetTaskListRequestObject) (tasks.GetTaskListResponseObject, error) {
	allTasks, err := t.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTaskList200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{Id: int64(tsk.ID), Title: tsk.Title, IsDone: tsk.IsDone}
		response = append(response, task)
	}

	return response, nil
}

func (t *TaskHandler) CreateTask(_ context.Context, request tasks.CreateTaskRequestObject) (tasks.CreateTaskResponseObject, error) {
	taskRequest := request.Body
	var response tasks.CreateTask201JSONResponse

	taskToCreate := tasksService.TaskCreateInput{Title: taskRequest.Title, IsDone: taskRequest.IsDone}

	createdTask, err := t.service.CreateTask(&taskToCreate)

	if err != nil {
		return nil, err
	}

	response = createdTask.ToResponse()

	return response, nil
}

func (t *TaskHandler) DeleteTask(_ context.Context, request tasks.DeleteTaskRequestObject) (tasks.DeleteTaskResponseObject, error) {
	id := request.Id
	err := t.service.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTask204Response{}, nil
}

func (t *TaskHandler) UpdateTask(_ context.Context, request tasks.UpdateTaskRequestObject) (tasks.UpdateTaskResponseObject, error) {
	requestBody := request.Body
	requestId := request.Id

	taskToUpdate := tasksService.TaskUpdateInput{
		Title:  requestBody.Title,
		IsDone: requestBody.IsDone,
	}

	updatedTask, err := t.service.UpdateTask(&taskToUpdate, requestId)
	if err != nil {
		return nil, err
	}

	response := tasks.UpdateTask200JSONResponse{
		Id:     int64(updatedTask.ID),
		Title:  updatedTask.Title,
		IsDone: updatedTask.IsDone,
	}

	return response, nil
}

func NewTaskHandler(service tasksService.TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}
