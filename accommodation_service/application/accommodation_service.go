package application

import (
	"booking-backend/accommodation_service/domain"
	pb "booking-backend/common/proto/accommodation_service"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) GetAll() ([]*pb.SingleAccommodation, error) {
	accommodations, err := service.store.GetAll()
	grpcAccommodations := ConvertToGrpcList(accommodations)
	return grpcAccommodations, err
}

func (service *AccommodationService) GetById(id string) (*pb.SingleAccommodation, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accommodation, err := service.store.GetById(objId)
	grpcAccommodation := ConvertToGrpc(accommodation)
	return grpcAccommodation, err
}

func (service *AccommodationService) Create(accommodation domain.Accommodation) error {
	return service.store.Create(accommodation)
}

func (service *AccommodationService) Update(accommodation domain.Accommodation) error {
	return service.store.Update(accommodation)
}

func (service *AccommodationService) Delete(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return service.store.Delete(objId)
}

func (service *AccommodationService) SearchAccommodations(request *pb.SearchAccommodationsRequest) ([]*pb.SingleAccommodation, error) {
	accommodations, err := service.store.GetAll()
	if err != nil {
		return nil, err
	}

	filteredAccommodations := make([]*pb.SingleAccommodation, 0)
	for _, accommodation := range accommodations {
		if accommodation.MinGuests <= request.GuestsNumber && accommodation.MaxGuests >= request.GuestsNumber && accommodation.Address.Country == request.Country && accommodation.Address.City == request.City {
			for _, appointment := range accommodation.Appointments {
				interval, _ := domain.StringToDateInterval(request.Interval)
				if isExactOverlap(appointment.Interval, interval) {
					filteredAccommodations = append(filteredAccommodations, ConvertToGrpc(&accommodation))
					break
				}
			}
		}
	}
	return filteredAccommodations, nil
}

func (service *AccommodationService) AddAppointment(appointment domain.CreateAppointment) error {
	accommodation, err := service.store.GetById(appointment.Id)
	if err != nil {
		return err
	}
	err = HasOverlap(appointment, accommodation.Appointments)
	if err != nil {
		return err
	}

	accommodation.Appointments = append(accommodation.Appointments, domain.MakeAppointmentFromCreateAppointment(appointment))
	return service.store.Update(*accommodation)
}

func (service *AccommodationService) UpdateAppointment(appointment domain.UpdateAppointment) error {
	accommodation, err := service.store.GetById(appointment.OldAppointment.Id)
	if err != nil {
		return err
	}
	appointments, err := RemoveOldAppointment(appointment.OldAppointment, accommodation.Appointments)
	if err != nil {
		return err
	}

	err = HasOverlap(appointment.NewAppointment, accommodation.Appointments)
	if err != nil {
		return err
	}
	accommodation.Appointments = append(appointments, domain.MakeAppointmentFromCreateAppointment(appointment.NewAppointment))
	return service.store.Update(*accommodation)
}

func (service *AccommodationService) ValidateReservation(accommodationId primitive.ObjectID, interval domain.DateInterval) error {
	accommodation, err := service.store.GetById(accommodationId)
	if err != nil {
		return err
	}

	return func(appointments []domain.Appointment, interval domain.DateInterval) error {
		for _, entity := range appointments {
			if isExactOverlap(interval, entity.Interval) {
				return nil
			}
		}
		return errors.New("No appointment with this interval exists")
	}(accommodation.Appointments, interval)
}

func RemoveOldAppointment(oldAppointment domain.CreateAppointment, appointments []domain.Appointment) ([]domain.Appointment, error) {
	for i, entity := range appointments {
		if isExactOverlap(oldAppointment.Interval, entity.Interval) {
			appointments[i] = appointments[len(appointments)-1]
			return appointments[:len(appointments)-1], nil
		}
	}
	return nil, errors.New("Old Appointment is not present in this Accommodation!")
}

func isExactOverlap(interval, accInterval domain.DateInterval) bool {
	return interval.DateFrom == accInterval.DateFrom && interval.DateTo == accInterval.DateTo
}

func HasOverlap(appointment domain.CreateAppointment, accAppointments []domain.Appointment) error {
	interval := appointment.Interval
	for _, entity := range accAppointments {
		err := checkOverlap(interval, entity.Interval)
		if err != nil {
			return err
		}
		err = checkOverlap(entity.Interval, interval)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkOverlap(interval domain.DateInterval, helperInterval domain.DateInterval) error {
	if helperInterval.DateFrom.Before(interval.DateFrom) && interval.DateFrom.Before(helperInterval.DateTo) {
		return fmt.Errorf("The Interval =>%+v --> %+v<= is Inside this Interval =>%+v --> %+v<=! INTERVAL OVERLAP!",
			interval.DateFrom, interval.DateTo, helperInterval.DateFrom, helperInterval.DateTo)
	}
	if helperInterval.DateFrom.Before(interval.DateTo) && interval.DateTo.Before(helperInterval.DateTo) {
		return fmt.Errorf("The Interval =>%+v --> %+v<= is Inside this Interval =>%+v --> %+v<=! INTERVAL OVERLAP!",
			interval.DateFrom, interval.DateTo, helperInterval.DateFrom, helperInterval.DateTo)
	}
	if interval.DateFrom == helperInterval.DateFrom && interval.DateTo == helperInterval.DateTo {
		return fmt.Errorf("The Intervals =>%+v --> %+v<= and Interval =>%+v --> %+v<= overlap exactly! EXACT OVERLAP!",
			interval.DateFrom, interval.DateTo, helperInterval.DateFrom, helperInterval.DateTo)
	}
	return nil
}

//GRPC CONVERTERS -> OVO IZMESTITI POSLE

func ConvertToGrpcList(accommodations []domain.Accommodation) []*pb.SingleAccommodation {
	var converted []*pb.SingleAccommodation

	for _, entity := range accommodations {
		newRes := ConvertToGrpc(&entity)
		converted = append(converted, newRes)
	}

	return converted
}

func ConvertToGrpc(accommodation *domain.Accommodation) *pb.SingleAccommodation {
	var allAppointments []*pb.AppointmentResponse

	for _, entity := range accommodation.Appointments {
		appointment := ConvertToGrpcAppointment(entity)
		allAppointments = append(allAppointments, appointment)
	}

	res := pb.SingleAccommodation{
		Id:        accommodation.Id.Hex(),
		Longitude: accommodation.Longitude,
		Latitude:  accommodation.Latitude,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Name:      accommodation.Name,
		Address: &pb.AccommodationAddress{
			City:    accommodation.Address.City,
			Number:  accommodation.Address.Number,
			Country: accommodation.Address.Country,
			Street:  accommodation.Address.Street,
		},
		FreeAppointments: allAppointments,
    Host: accommodation.Host,
	}

	return &res
}

func ConvertToGrpcAppointment(appointment domain.Appointment) *pb.AppointmentResponse {
	interval := ConvertToGrpcInterval(appointment.Interval)

	return &pb.AppointmentResponse{
		Interval: interval,
		Price:    appointment.Price,
	}
}

func ConvertToGrpcInterval(interval domain.DateInterval) *pb.SingleDateInterval {

	return &pb.SingleDateInterval{
		DateFrom: interval.DateFrom.Format(time.DateOnly),
		DateTo:   interval.DateTo.Format(time.DateOnly),
	}
}
