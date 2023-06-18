package application

import (
	"context"
	"errors"
	"fmt"
	"log"

	"booking-backend/common/clients"
	"booking-backend/common/messaging"
	"booking-backend/common/proto/notification_service"
	"booking-backend/common/proto/reservation_service"
	pb "booking-backend/common/proto/user_service"
	"booking-backend/user-service/domain"
	"booking-backend/user-service/startup/config"
)

type UserService struct {
<<<<<<< HEAD
	store       domain.UserStore
	ratingStore domain.RatingStore
  subscriber  messaging.SubscriberModel
  notificationPublisher messaging.PublisherModel
}

func NewUserService(store domain.UserStore, ratingStore domain.RatingStore, subscriber messaging.SubscriberModel, notificationPublisher messaging.PublisherModel) (*UserService, error) {
  service := &UserService{
		store:       store,
		ratingStore: ratingStore,
    subscriber: subscriber,
    notificationPublisher: notificationPublisher,
=======
	store      domain.UserStore
	subscriber messaging.SubscriberModel
}

func NewUserService(store domain.UserStore, subscriber messaging.SubscriberModel) (*UserService, error) {
	service := &UserService{
		store:      store,
		subscriber: subscriber,
>>>>>>> 3be3f21 (Implemented saga)
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
	err = service.store.Create(user)

  if err == nil {
    notifications := getNotificationClient()
    _, err = notifications.NewUserSettings(context.Background(),&notification_service.NewUserSettingsRequest{
      Host: user.Email,
      Role: user.Role,
    })
  }

  return err
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

func (service *UserService) ProminentUser(host string) {
	reservation := getReservationClient()

<<<<<<< HEAD
  user, err:= service.GetOne(host)
  if err != nil {
    return
  }

  response, err := reservation.GetHostAnalytics(context.Background(), &reservation_service.HostAnalyticsRequest{Host: host})
=======
	response, err := reservation.GetHostAnalytics(context.Background(), &reservation_service.HostAnalyticsRequest{Host: host})
>>>>>>> 3be3f21 (Implemented saga)

	if err != nil {
		return
	}

<<<<<<< HEAD
  prominent := isProminent(response)
  if (user.IsProminent != prominent) {
    service.store.UpdateProminent(prominent, host)

    content := "You lost your prominent status"
    if prominent {
        content = "You gained the prominent status, congrats!"
    }

    service.notificationPublisher.Publish(messaging.NotificationMessage{
      User: host,
      Subject: "Your prominent status has changed!",
      Content: content,
      Type: messaging.ReservationRequest,
    })
  }

=======
	service.store.UpdateProminent(isProminent(response), host)
>>>>>>> 3be3f21 (Implemented saga)
}

func isProminent(reservationAnalytics *reservation_service.HostAnalyticsResponse) bool {
	return reservationAnalytics.CancelRate < 5.0 &&
		reservationAnalytics.ExpiredCount >= 5 &&
		reservationAnalytics.IntervalCount > 50
}

func getReservationClient() reservation_service.ReservationServiceClient {
	return clients.NewReservationClient(fmt.Sprintf("%s:%s", config.NewConfig().ReservationHost, config.NewConfig().ReservationPort))
}

func getNotificationClient() notification_service.NotificationServiceClient {
	return clients.NewNotificationClient(fmt.Sprintf("%s:%s", config.NewConfig().NotificationHost, config.NewConfig().NotificationPort))
}
