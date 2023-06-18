package domain

import (
	pb "booking-backend/common/proto/rating_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RatingAccommodation struct {
	Accommodation string `json:"accommodation" bson:"accommodation"`
	User          string `json:"user" bson:"user"`
	Date          string `json:"date" bson:"date"`
	Rate          int32  `json:"rate" bson:"rate"`
}

type RatingUser struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	RatedUser string             `json:"ratedUser" bson:"ratedUser"`
	RatedBy   string             `json:"ratedBy" bson:"ratedBy""`
	Date      string             `json:"date" bson:"date"`
	Rate      int32              `json:"rate" bson:"rate"`
	Status    Status             `json:"status" bson:"status"`
}

type Status int8

const (
	Pending Status = iota
	Approved
	Canceled
)

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

func UserRatingsToGrpcList(ratings []RatingUser) []*pb.UserRating {
	var converted []*pb.UserRating
	for _, entity := range ratings {
		newRes := UserRatingToGrpcRate(entity)
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

func UserRatingToGrpcRate(rating RatingUser) pb.UserRating {
	return pb.UserRating{
		RatedUser: rating.RatedUser,
		RatedBy:   rating.RatedBy,
		Rate:      rating.Rate,
		Date:      rating.Date,
	}
}

func MakeRating(rating *pb.RateUserRequest) RatingUser {
	return RatingUser{
		Rate:      rating.Rate,
		RatedBy:   rating.RatedBy,
		RatedUser: rating.RatedUser,
		Date:      time.Now().Format("2006-01-02"),
	}
}
