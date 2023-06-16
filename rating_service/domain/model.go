package domain

import (
	pb "booking-backend/common/proto/rating_service"
	"time"
)

type RatingAccommodation struct {
	Accommodation string `json:"accommodation" bson:"accommodation"`
	User          string `json:"user" bson:"user"`
	Date          string `json:"date" bson:"date"`
	Rate          int32  `json:"rate" bson:"rate"`
}

func MakeRate(rate int32, accommodation, user string) RatingAccommodation {
	return RatingAccommodation{
		Accommodation: accommodation,
		User:          user,
		Date:          time.Now().Format("2006-01-02"),
		Rate:          rate,
	}
}

func RateListToGrpcRateList(rates []RatingAccommodation) []*pb.RateAccommodationResponse {
	var converted []*pb.RateAccommodationResponse

	for _, entity := range rates {
		newRes := RateToGrpcRate(entity)
		converted = append(converted, &newRes)
	}

	return converted
}

func RateToGrpcRate(rate RatingAccommodation) pb.RateAccommodationResponse {
	return pb.RateAccommodationResponse{
		User: rate.User,
		Date: rate.Date,
		Rate: rate.Rate,
	}
}
