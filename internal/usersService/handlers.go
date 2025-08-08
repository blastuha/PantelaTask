package usersService

import (
	"context"
	"errors"
	"fmt"
	"task1/internal/utils"

	api "task1/internal/web/users"
)

type usersHandler struct {
	svc UsersService
}

func NewUsersHandler(svc UsersService) api.StrictServerInterface {
	return &usersHandler{svc: svc}
}

func (h *usersHandler) GetTasksForUser(
	_ context.Context,
	request api.GetTasksForUserRequestObject,
) (api.GetTasksForUserResponseObject, error) {
	rawID := request.Id

	uid64, err := utils.ParseUintID(rawID)
	if err != nil {
		return api.GetTasksForUser400JSONResponse{Error: "invalid user ID"}, nil
	}

	userID := uint(uid64)

	tasks, err := h.svc.GetTasksForUser(userID)
	if err != nil {
		if errors.Is(err, ErrUserNoFound) {
			return api.GetTasksForUser404JSONResponse{Error: "user not found"}, nil
		}
		return nil, err
	}

	resp := make(api.GetTasksForUser200JSONResponse, len(tasks))
	for i, t := range tasks {
		resp[i] = api.Task{Id: uint32(t.ID), Title: t.Title, IsDone: t.IsDone}
	}
	return resp, nil
}

func (h *usersHandler) GetAllUsers(_ context.Context, _ api.GetAllUsersRequestObject) (api.GetAllUsersResponseObject, error) {
	usersList, err := h.svc.GetAllUsers()
	if err != nil {
		return api.GetAllUsers500JSONResponse{Error: "internal server error"}, nil
	}

	resp := make([]api.UserResponse, len(usersList))
	for i := range usersList {
		resp[i] = toUserResponse(&usersList[i])
	}
	return api.GetAllUsers200JSONResponse(resp), nil
}

func (h *usersHandler) CreateUser(_ context.Context, req api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	user, err := h.svc.CreateUser(*req.Body)
	if err != nil {
		return api.CreateUser400JSONResponse{Error: err.Error()}, nil
	}
	return api.CreateUser201JSONResponse(toUserResponse(user)), nil
}

func (h *usersHandler) UpdateUser(_ context.Context, req api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	updatedUser, err := h.svc.UpdateUser(req.Id, *req.Body)
	if errors.Is(err, ErrUserNoFound) {
		return api.UpdateUser404JSONResponse{Error: "user not found"}, nil
	}
	if err != nil {
		return api.UpdateUser400JSONResponse{Error: err.Error()}, nil
	}
	return api.UpdateUser200JSONResponse(toUserResponse(updatedUser)), nil
}

func (h *usersHandler) DeleteUser(_ context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	err := h.svc.DeleteUser(request.Id)
	if errors.Is(err, ErrUserNoFound) {
		return api.DeleteUser404JSONResponse{
			Error: "user not found",
		}, nil
	}
	if err != nil {
		return api.DeleteUser500JSONResponse{
			Error: fmt.Sprintf("internal error: %v", err),
		}, nil
	}

	return api.DeleteUser204Response{}, nil
}
