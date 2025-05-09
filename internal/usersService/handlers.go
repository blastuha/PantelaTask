package usersService

import (
	"context"
	"errors"
	"fmt"

	api "task1/internal/web/users"
)

type usersHandler struct {
	svc UsersService
}

func NewUsersHandler(svc UsersService) api.StrictServerInterface {
	return &usersHandler{svc: svc}
}

func toUserResponse(u *User) api.UserResponse {
	if u == nil {
		return api.UserResponse{}
	}
	id := int64(u.ID)
	email := u.Email
	return api.UserResponse{
		Id:    &id,
		Email: &email,
	}
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

func (h *usersHandler) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
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
