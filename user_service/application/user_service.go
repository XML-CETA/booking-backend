package application

import (
	"errors"

	pb "booking-backend/common/proto/user_service"
	"booking-backend/user-service/domain"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Create(user domain.User) error {
	_, err := service.store.GetOne(user.Email)
	if err == nil {
		return errors.New("User with this username already exists")
	}
	return service.store.Create(user)
}

func (service *UserService) Delete(username string) error {
	return service.store.Delete(username)
}

func (service *UserService) Update(user domain.User) error {
	return service.store.Update(user)
}

func (service *UserService) GetOne(email string) (domain.User, error) {
	return service.store.GetOne(email)
}

func (service *UserService) LoginCheck(email string, password string) (pb.User, error) {
	user, err := service.store.GetOne(email)
	userRPC := service.UserToRPC(user)
	if err != nil {
		return userRPC, err
	}
	if user.Password != password {
		return userRPC, errors.New("Incorect password")
	}
	return userRPC, nil
}

func (service *UserService) UserToRPC(user domain.User) pb.User {
	rpcUser := pb.User{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Surname:  user.Surname,
		Address: &pb.Address{
			City:    user.Address.City,
			Number:  user.Address.Number,
			Country: user.Address.Country,
			Street:  user.Address.Street,
		},
		Role: user.Role,
	}
	return rpcUser
}
