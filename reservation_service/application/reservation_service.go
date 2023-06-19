package application

import (
	"booking-backend/common/clients"
	"booking-backend/common/messaging"
	"booking-backend/common/proto/accommodation_service"
	pb "booking-backend/common/proto/reservation_service"
	"booking-backend/reservation-service/domain"
	"booking-backend/reservation-service/startup/config"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store                  domain.ReservationStore
	prominentHostPublisher messaging.PublisherModel
	notificationPublisher  messaging.PublisherModel
}

func NewReservationService(store domain.ReservationStore, prominentHostPublisher messaging.PublisherModel, notificationPublisher messaging.PublisherModel) *ReservationService {
	return &ReservationService{
		store:                  store,
		prominentHostPublisher: prominentHostPublisher,
		notificationPublisher:  notificationPublisher,
	}
}

func (service *ReservationService) CreateReservation(reservation domain.Reservation) error {
	_, err := service.store.GetFirstActive(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		return errors.New("Could not create, an active reservation with the same interval already exists")
	}

	response, err := validateReservation(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	reservation.Host = response.Host
	isAutomatic, err := isAutomaticConfirmation(reservation.Accommodation)

	if err != nil {
		return err
	}

	if isAutomatic.IsAutomaticConfirmation {
		reservation.Status = domain.Reserved
	}

	reservation.Duration, err = calculateDuration(reservation.DateFrom, reservation.DateTo)

	if err != nil {
		return err
	}

	err = service.store.CreateReservation(reservation)

	if err == nil {
		service.prominentHostPublisher.Publish(reservation.Host)

		service.notificationPublisher.Publish(messaging.NotificationMessage{
			User:    reservation.Host,
			Subject: "You have a new reservation request!",
			Content: fmt.Sprintf("%v requested for a reservation on accommodation with id: %v", reservation.User, reservation.Accommodation),
			Type:    messaging.ReservationRequest,
		})
	}

	return err
}

func (service *ReservationService) IsAppointmentReserved(accommodation, dateFrom, dateTo string) (bool, error) {
	_, err := service.store.GetFirstActive(accommodation, dateFrom, dateTo)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (service *ReservationService) GetAll() ([]domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) GetWaitingReservations(host string) ([]domain.Reservation, error) {
	return service.store.GetWaitingReservations(host)
}

func (service *ReservationService) ConfirmReservation(reservationId string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}
	reservation, err := service.store.GetById(id)
	err = service.store.ConfirmReservation(id)
	if err != nil {
		return err
	}

	err = service.cancelReservationsWithOverlap(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		service.prominentHostPublisher.Publish(reservation.Host)

		service.notificationPublisher.Publish(messaging.NotificationMessage{
			User:    reservation.User,
			Subject: "Your reservation has been approved!",
			Content: fmt.Sprintf("Reservation lasting from %v to %v has been approved!", reservation.DateFrom, reservation.DateTo),
			Type:    messaging.ReservationResponse,
		})
	}

	return err
}

func (service *ReservationService) CheckIfUserVisitedHost(user, host string) (bool, error) {
	reservations, err := service.store.GetByUserAndHost(user, host)
	if err != nil {
		return false, err
	}
	return len(reservations) >= 1, nil
}

func (service *ReservationService) Delete(reservationId string, user string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}

	reservation, err := service.store.GetByIdAndUser(id, user)

	if reservation.Status == domain.Waiting {
		return service.store.Cancel(id)
	}

	if !canDeleteReservation(reservation) {
		return errors.New("Can not delete, reservation starts in less then a day or is in progress")
	}

	err = service.store.Cancel(id)

	if err == nil {
		service.prominentHostPublisher.Publish(reservation.Host)

		service.notificationPublisher.Publish(messaging.NotificationMessage{
			User:    reservation.Host,
			Subject: "A reservation has been canceled!",
			Content: fmt.Sprintf("Reservation lasting from %v to %v has been canceled :(", reservation.DateFrom, reservation.DateTo),
			Type:    messaging.ReservationCancel,
		})
	}

	return err
}

func (service *ReservationService) Decline(reservationId string) error {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return err
	}

	return service.store.Decline(id)
}

func (service *ReservationService) ConvertToGrpcList(reservations []domain.Reservation) []*pb.Reservation {
	var converted []*pb.Reservation

	for _, entity := range reservations {
		newRes := pb.Reservation{
			Id:            entity.Id.Hex(),
			Accommodation: entity.Accommodation,
			DateFrom:      entity.DateFrom,
			DateTo:        entity.DateTo,
			Guests:        entity.Guests,
			User:          entity.User,
			Status:        int32(entity.Status),
			Host:          entity.Host,
		}

		converted = append(converted, &newRes)
	}

	return converted
}

func (service *ReservationService) ConvertToGrpcWaitingReservations(reservations []domain.Reservation) []*pb.WaitingReservation {
	var converted []*pb.WaitingReservation

	for _, entity := range reservations {
		canceledCount, _ := service.store.CountUserCanceled(entity.User)
		newRes := pb.WaitingReservation{
			Id:                       entity.Id.Hex(),
			Accommodation:            entity.Accommodation,
			DateFrom:                 entity.DateFrom,
			DateTo:                   entity.DateTo,
			Guests:                   entity.Guests,
			User:                     entity.User,
			UserCanceledReservations: canceledCount,
		}
		converted = append(converted, &newRes)
	}
	return converted
}

func (service *ReservationService) cancelReservationsWithOverlap(accommodationId string, dateFrom, dateTo string) error {
	reservations, err := service.store.GetWaitingByAccommodation(accommodationId)
	if err != nil {
		return err
	}
	for _, entity := range reservations {
		if isExactOverlap(entity, dateFrom, dateTo) {
			service.store.Cancel(entity.Id)
		}
	}
	return nil
}

func (service *ReservationService) GetCancelRate(host string) (float32, error) {
	nonCanceled, err := service.store.CountNonCanceled(host)
	if err != nil {
		return 0.0, err
	}

	canceled, err := service.store.CountCanceled(host)
	if err != nil || canceled == 0 {
		return 0.0, err
	}

	return float32(nonCanceled / canceled), nil
}

func (service *ReservationService) GetExpiredCount(host string) (int32, error) {
	return service.store.CountExpired(host)
}

func (service *ReservationService) GetIntervalCount(host string) (int32, error) {
	return service.store.GetHostIntervalSum(host)
}

func (service *ReservationService) HasActiveReservations(user string, role string) (bool, error) {
	count, err := service.store.CountActive(user, role)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func getAccommodationClient() accommodation_service.AccommodationServiceClient {
	return clients.NewAccommodationClient(fmt.Sprintf("%s:%s", config.NewConfig().AccommodationServiceHost, config.NewConfig().AccommodationServicePort))
}

func validateReservation(accommodationId, dateFrom, dateTo string) (*accommodation_service.ValidateReservationResponse, error) {
	accommodation := getAccommodationClient()

	return accommodation.ValidateReservation(context.Background(), &accommodation_service.ValidateReservationRequest{
		Accommodation: accommodationId,
		DateFrom:      dateFrom,
		DateTo:        dateTo,
	})
}

func isAutomaticConfirmation(accommodationId string) (*accommodation_service.IsAutomaticConfirmationResponse, error) {
	accommodation := getAccommodationClient()
	return accommodation.IsAutomaticConfirmation(context.Background(), &accommodation_service.AccommodationIdRequest{
		Id: accommodationId,
	})
}

func isExactOverlap(reservation domain.Reservation, dateFrom, dateTo string) bool {
	return reservation.DateFrom == dateFrom && reservation.DateTo == dateTo
}

func checkReservationDate(dateFrom string) bool {
	date, err := time.Parse(time.DateOnly, dateFrom)
	if err != nil {
		return false
	}

	now := time.Now().UTC()
	return now.UTC().Compare(date.AddDate(0, 0, -1)) == -1
}

func canDeleteReservation(reservation domain.Reservation) bool {
	return reservation.Status == domain.Reserved && checkReservationDate(reservation.DateFrom)
}

func calculateDuration(dateFromString, dateToString string) (int32, error) {
	dateFrom, err := time.Parse(time.DateOnly, dateFromString)
	if err != nil {
		return 0, err
	}

	dateTo, err := time.Parse(time.DateOnly, dateToString)
	if err != nil {
		return 0, err
	}

	difference := dateTo.Sub(dateFrom)

	return int32(difference.Hours() / 24), nil
}
