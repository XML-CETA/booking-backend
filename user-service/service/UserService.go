package service

import (
	"example/grpc/model"
	"example/grpc/repo"
)

type UserService struct {
	Repo *repo.UserRepository
}

func (service *UserService) Create(user *model.User) error {
	if(isUsernameUnique(user.Username)){
		return service.Repo.Create(user)
	}else 
		 error("username taken")
}

func (service *UserService) Delete(id primitive.ObjectID) error {
	return service.Repo.Delete(id)
}

func (service *UserService) Update(flight *model.Flight) error {
	return service.Repo.Update(flight)
}
