package service

import (
	"errors"
	"example/grpc/model"
	"example/grpc/proto/users"
	"example/grpc/repo"
)

type UserService struct {
	Repo *repo.UserRepository
}

func (service *UserService) Create(user *model.User) error {
	_, err := service.Repo.GetOne(user.Email)
	if err == nil {
		return errors.New("User with this username already exists")
	}
	return service.Repo.Create(user)
}

func (service *UserService) Delete(username string) error {
	return service.Repo.Delete(username)
}

func (service *UserService) Update(user *model.User) error {
	return service.Repo.Update(user)
}

func (service *UserService) GetOne(email string) (model.User, error) {
	return service.Repo.GetOne(email)
}

func (service *UserService) LoginCheck(email string, password string) (users.User, error) {
	user, err := service.Repo.GetOne(email)
	userRPC := service.UserToRPC(user)
	if err != nil {
		return userRPC, err
	}
	if user.Password != password {
		return userRPC, errors.New("Incorect password")
	}
	return userRPC, nil
}

func (service *UserService) UserToRPC(user model.User) users.User {
	rpcUser := users.User{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Surname:  user.Surname,
		Address: &users.Address{
			City:    user.Address.City,
			Number:  user.Address.Number,
			Country: user.Address.Country,
			Street:  user.Address.Street,
		},
		Role: user.Role,
	}
	return rpcUser
}
