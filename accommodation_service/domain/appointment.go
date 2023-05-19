package domain

import (
	pb "booking-backend/common/proto/accommodation_service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	Interval DateInterval `json:"interval"  bson:"interval"`
	Price    float64      `json:"price"  bson:"price"`
}

type CreateAppointment struct {
	Id       primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	Interval DateInterval       `json:"interval"  bson:"interval"`
	Price    float64            `json:"price"  bson:"price"`
}

type UpdateAppointment struct {
	OldAppointment CreateAppointment `json:"oldAppointment"  bson:"oldAppointment"`
	NewAppointment CreateAppointment `json:"newAppointment"  bson:"newAppointment"`
}

type DateInterval struct {
	DateFrom time.Time `json:"dateFrom"  bson:"dateFrom"`
	DateTo   time.Time `json:"dateTo"  bson:"dateTo"`
}

func MakeAppointmentFromCreateAppointment(appointment CreateAppointment) Appointment {
	return Appointment{
		Interval: appointment.Interval,
		Price:    appointment.Price,
	}
}

func MakeUpdateAppointment(appointments *pb.UpdateAppointmentRequest) (UpdateAppointment, error) {
	oldAppointment, err := MakeCreateAppointment(appointments.OldAppointment)
	if err != nil {
		return UpdateAppointment{}, err
	}

	newAppointment, err := MakeCreateAppointment(appointments.NewAppointment)
	if err != nil {
		return UpdateAppointment{}, err
	}

	return UpdateAppointment{
		NewAppointment: newAppointment,
		OldAppointment: oldAppointment,
	}, nil
}

func MakeCreateAppointment(appointment *pb.SingleAppointment) (CreateAppointment, error) {
	id, err := primitive.ObjectIDFromHex(appointment.AccommodationId)

	if err != nil {
		return CreateAppointment{}, err
	}

	interval, err := StringToDateInterval(appointment.Interval)

	if err != nil {
		return CreateAppointment{}, err
	}

	return CreateAppointment{
		Id:       id,
		Interval: interval,
		Price:    appointment.Price,
	}, nil

}

func StringToDateInterval(interval *pb.SingleDateInterval) (DateInterval, error) {
	dateFrom, err := time.Parse(time.DateOnly, interval.DateFrom)
	if err != nil {
		return DateInterval{}, err
	}

	dateTo, err := time.Parse(time.DateOnly, interval.DateTo)
	if err != nil {
		return DateInterval{}, err
	}

	return DateInterval{
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}, nil
}
