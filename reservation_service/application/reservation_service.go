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
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store                  domain.ReservationStore
	prominentHostPublisher messaging.PublisherModel
}

func NewReservationService(store domain.ReservationStore, prominentHostPublisher messaging.PublisherModel) *ReservationService {
	return &ReservationService{
		store:                  store,
		prominentHostPublisher: prominentHostPublisher,
	}
}

func (service *ReservationService) CreateReservation(reservation domain.Reservation) error {
	fmt.Println("USO U SERVICE")

	_, err := service.store.GetFirstActive(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	if err == nil {
		return errors.New("Could not create, an active reservation with the same interval already exists")
	}

	fmt.Println("PROSO GET FIRST ACTIVE")

	response, err := validateReservation(reservation.Accommodation, reservation.DateFrom, reservation.DateTo)

	fmt.Println(response.Success)
	fmt.Println(response.Host)

	fmt.Println("NASO")

	reservation.Host = response.Host
	isAutomatic, err := isAutomaticConfirmation(reservation.Accommodation)

	if err != nil {
		return err
	}

	fmt.Println(isAutomatic)

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
	}

	return err
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
	}

	return err
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
	}

	return err
}

func (service *ReservationService) GetReservationFlights(reservationId, city, user string, isArrival bool) (*pb.FlightForReservationResponse, error) {
	id, err := primitive.ObjectIDFromHex(reservationId)
	if err != nil {
		return nil, err
	}

	fmt.Println("PROSO ID")
	fmt.Println(id)

	reservation, err := service.store.GetByIdAndUser(id, user)
	if err != nil {
		return nil, err
	}

	fmt.Println("PROSO RESERVATION")
	fmt.Println(reservation.Id)

	accommodation, err := getAccommodation(reservation.Accommodation)
	if err != nil {
		return nil, err
	}

	fmt.Println("PROSO ACCOMMODATION")
	fmt.Println(accommodation.Id)

	var flights []domain.Flight
	if isArrival {
		flights, err = GetReservationFlights(city, accommodation.Address.City, reservation.DateFrom)
	} else {
		flights, err = GetReservationFlights(accommodation.Address.City, city, reservation.DateTo)
	}

	hasAccount := hasAirlineAccount(user)

	return &pb.FlightForReservationResponse{
		HasAirlineAccount: hasAccount,
		Flights:           ConvertFlightsToGrpcList(flights),
	}, err
}

func GetReservationFlights(startingPoint, destination, date string) ([]domain.Flight, error) {
	fmt.Println("USO U GRPC F-JU")

	// serverPort := 3000

	requestURL := fmt.Sprintf("http://172.21.0.4:3000/flights/reservation/")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	return nil, nil
}

func hasAirlineAccount(user string) bool {
	return false
}

func getAccommodation(accommodationId string) (*accommodation_service.SingleAccommodation, error) {
	accommodation := getAccommodationClient()

	return accommodation.GetById(context.Background(), &accommodation_service.AccommodationIdRequest{
		Id: accommodationId,
	})
}

func ConvertFlightsToGrpcList(flights []domain.Flight) []*pb.Flight {
	var converted []*pb.Flight

	for _, entity := range flights {
		newRes := pb.Flight{
			Id:             entity.Id.Hex(),
			DateAndTime:    entity.FlighDateAndTime.Format("2006-01-02 15:04:05"),
			StartingPoint:  entity.StartingPoint,
			Destination:    entity.Destination,
			Price:          entity.Price,
			RemainingSeats: entity.RemainingSeats,
		}

		converted = append(converted, &newRes)
	}
	return converted
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
