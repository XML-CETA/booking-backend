package api

import (
	"context"
	"errors"
	"fmt"

	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	pb "booking-backend/common/proto/user_service"
	"booking-backend/user-service/application"
	"booking-backend/user-service/domain"
	"booking-backend/user-service/startup/config"

	"google.golang.org/grpc/metadata"
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
  user, err := Authorize(ctx, []string{"HOST", "REGULAR"})
  if err != nil {
    return nil, err
  }

  userData, err := h.service.GetOne(user)
  if err != nil {
    return nil, err
  }

  canDelete, err := h.service.CanDelete(user, userData.Role)
  if err != nil {
    return nil, err
  }

  if !canDelete {
    return nil, errors.New("Unable to delete, user has leftover reservations")
  }

  if userData.Role == "HOST" {
    err = h.service.DeleteHostAccommodations(user)
    if err != nil {
      return nil, err
    }
  }

	err = h.service.Delete(user)
	if err != nil {
    return nil, err
	}

	return &pb.DeleteResponse{
		Message: fmt.Sprintf("Succesfully deleted!"),
	}, nil
}

func (h UserHandler) GetUserData(ctx context.Context, request *pb.GetRequest) (*pb.UserFull, error) {
  user, err := Authorize(ctx, []string{"HOST", "REGULAR"})
  if err != nil {
    return nil, err
  }

  userData, err := h.service.GetOne(user)

  if err != nil {
    return nil, err
  }

	return &pb.UserFull{
    Email: userData.Email,
    Password: userData.Password,
    Name: userData.Name,
    Surname: userData.Surname,
    Role: userData.Role,
    IsProminent: userData.IsProminent,
    Address: &pb.Address{
      City: userData.Address.City,
      Country: userData.Address.Country,
      Street: userData.Address.Street,
      Number: userData.Address.Number,
    },
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

func Authorize(ctx context.Context, roleGuard []string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

  if err != nil {
    return "", err
  }

	return user.UserEmail, nil
}
