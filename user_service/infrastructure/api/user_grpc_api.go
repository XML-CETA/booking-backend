package api

import (
	"context"
	"fmt"

	pb "booking-backend/common/proto/user_service"
	"booking-backend/user-service/application"
	"booking-backend/user-service/domain"
)

type UserHandler struct {
	pb.UnimplementedUsersServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h UserHandler) RateUser(ctx context.Context, request *pb.RateUserRequest) (*pb.RateUserResponse, error) {
	rating := domain.MakeRating(request.Rating.Rating, request.Rating.RatedBy, request.Rating.RatedHost)
	err := h.service.RateUser(rating)
	if err != nil {
		return &pb.RateUserResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.RateUserResponse{
		Message: fmt.Sprintf("User is succesfully rated"),
	}, nil
}

func (h UserHandler) CreateUser(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	reqUser := request.User
	u := domain.User{
		Name:     reqUser.Name,
		Surname:  reqUser.Surname,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Address: domain.Address{
			Street:  reqUser.Address.Street,
			Number:  reqUser.Address.Number,
			City:    reqUser.Address.City,
			Country: reqUser.Address.Country,
		},
		Role: reqUser.Role,
	}

	err := h.service.Create(u)
	if err != nil {
		return &pb.CreateResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.CreateResponse{
		Message: fmt.Sprintf("Succesfully created"),
	}, nil
}

func (h UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	reqUser := request.User
	u := domain.User{
		Name:     reqUser.Name,
		Surname:  reqUser.Surname,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Address: domain.Address{
			Street:  reqUser.Address.Street,
			Number:  reqUser.Address.Number,
			City:    reqUser.Address.City,
			Country: reqUser.Address.Country,
		},
		Role: reqUser.Role,
	}

	err := h.service.Update(u)
	if err != nil {
		return &pb.UpdateResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.UpdateResponse{
		Message: fmt.Sprintf("Succesfully updated"),
	}, nil
}

func (h UserHandler) DeleteUser(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := h.service.Delete(request.UserEmail)
	if err != nil {
		return &pb.DeleteResponse{
			Message: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.DeleteResponse{
		Message: fmt.Sprintf("Succesfully deleted!"),
	}, nil
}

func (h UserHandler) LoginCheck(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.service.LoginCheck(request.Email, request.Password)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		User: &user,
	}, nil
}
