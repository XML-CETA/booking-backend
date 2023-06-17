package application

import (
	"context"
	"errors"
	"fmt"
	"log"

	"booking-backend/common/clients"
	"booking-backend/common/messaging"
	"booking-backend/common/proto/reservation_service"
	pb "booking-backend/common/proto/user_service"
	"booking-backend/user-service/domain"
	"booking-backend/user-service/startup/config"
)

type UserService struct {
	store       domain.UserStore
	ratingStore domain.RatingStore
  subscriber  messaging.SubscriberModel
}

func NewUserService(store domain.UserStore, ratingStore domain.RatingStore, subscriber messaging.SubscriberModel) (*UserService, error) {
  service := &UserService{
		store:       store,
		ratingStore: ratingStore,
    subscriber: subscriber,
	}

  err := service.subscriber.Subscribe(service.ProminentUser)

  log.Printf("%v", err)

  if err != nil {
    return nil, err
  }

  return service, nil
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

func (service *UserService) RateUser(rating domain.Rating) error {
	return service.ratingStore.Create(rating)
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

func (service *UserService) ProminentUser(host string) {
  reservation := getReservationClient()

  response, err := reservation.GetHostAnalytics(context.Background(), &reservation_service.HostAnalyticsRequest{Host: host})

  if err != nil {
    return
  }

  service.store.UpdateProminent(isProminent(response), host)
}

func isProminent(reservationAnalytics *reservation_service.HostAnalyticsResponse) bool {
  return reservationAnalytics.CancelRate < 5.0 &&
        reservationAnalytics.ExpiredCount >= 5 &&
        reservationAnalytics.IntervalCount > 50
}

func getReservationClient() reservation_service.ReservationServiceClient {
	return clients.NewReservationClient(fmt.Sprintf("%s:%s", config.NewConfig().ReservationHost, config.NewConfig().ReservationPort))
}
