package handlers

import (
	"context"
	"example/grpc/proto/users"
	"fmt"
)

type UserHandler struct {
	
}

func (h UserHandler) CreateUser(ctx context.Context, request *users.CreateRequest) (*users.CreateResponse, error) {
	return &users.CreateResponse{
		Error:,
	}, nil
}

func (h UserHandler) UpdateUser(ctx context.Context, request *users.UpdateRequest) (*users.UpdateResponse, error) {
	return &users.UpdateResponse{
		Error: fmt.Sprintf("Hi get!"),
	}, nil
}

func (h UserHandler) DeleteUser(ctx context.Context, request *users.DeleteRequest) (*users.DeleteResponse, error) {
	return &users.DeleteResponse{
		Error: fmt.Sprintf("Hi get!"),
	}, nil
}
