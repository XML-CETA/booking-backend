package handlers

import (
	"context"
	"example/grpc/model"
	"example/grpc/proto/users"
	"example/grpc/service"
	"fmt"
)

type UserHandler struct {
	users.UnimplementedUsersServiceServer
	Service *service.UserService
}

func (h UserHandler) CreateUser(ctx context.Context, request *users.CreateRequest) (*users.CreateResponse, error) {
	reqUser := request.User
	u := &model.User{
		Name:     reqUser.Name,
		Surname:  reqUser.Surname,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Address: model.Address{
			Street:  reqUser.Address.Street,
			Number:  reqUser.Address.Number,
			City:    reqUser.Address.City,
			Country: reqUser.Address.Country,
		},
		Role: reqUser.Role,
	}

	err := h.Service.Create(u)
	if err != nil {
		return &users.CreateResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &users.CreateResponse{
		Message: fmt.Sprintf("Succesfully created"),
	}, nil
}

func (h UserHandler) UpdateUser(ctx context.Context, request *users.UpdateRequest) (*users.UpdateResponse, error) {
	reqUser := request.User
	u := &model.User{
		Name:     reqUser.Name,
		Surname:  reqUser.Surname,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Address: model.Address{
			Street:  reqUser.Address.Street,
			Number:  reqUser.Address.Number,
			City:    reqUser.Address.City,
			Country: reqUser.Address.Country,
		},
		Role: reqUser.Role,
	}

	err := h.Service.Update(u)
	if err != nil {
		return &users.UpdateResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &users.UpdateResponse{
		Message: fmt.Sprintf("Succesfully updated"),
	}, nil
}

func (h UserHandler) DeleteUser(ctx context.Context, request *users.DeleteRequest) (*users.DeleteResponse, error) {
	err := h.Service.Delete(request.UserEmail)
	if err != nil {
		return &users.DeleteResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &users.DeleteResponse{
		Message: fmt.Sprintf("Succesfully deleted!"),
	}, nil
}

func (h UserHandler) LoginCheck(ctx context.Context, request *users.LoginRequest) (*users.LoginResponse, error) {
	user, err := h.Service.LoginCheck(request.Email, request.Password)

	if err != nil {
		return nil, err
	}

	return &users.LoginResponse{
		User: &user,
	}, nil
}
