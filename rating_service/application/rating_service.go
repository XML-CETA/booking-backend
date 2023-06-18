package application

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/accommodation_service"
	pb "booking-backend/common/proto/rating_service"
	"booking-backend/rating-service/domain"
	"booking-backend/rating-service/startup/config"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingService struct {
	rateAccommodationStore domain.RatingAccommodationStore
	rateUserStore          domain.RatingUserStore
	orchestrator           *RateUserOrchestrator
}

func NewRatingService(accommodationStore domain.RatingAccommodationStore, ratingUserStore domain.RatingUserStore, orchestrator *RateUserOrchestrator) *RatingService {
	return &RatingService{
		rateAccommodationStore: accommodationStore,
		rateUserStore:          ratingUserStore,
		orchestrator:           orchestrator,
	}
}

func (service *RatingService) CreateAccommodationRate(newRate domain.RatingAccommodation) error {
	_, err := accommodationExists(newRate.Accommodation)
	if err == nil {
		return errors.New("Given accommodation doesn't exist!")
	}

	if newRate.Rate < 1 || newRate.Rate > 5 {
		return errors.New("The rating should be between 1 and 5!")
	}

	_, err = service.rateAccommodationStore.GetByUserAndAccommodation(newRate.User, newRate.Accommodation)
	if err == nil {
		return errors.New("Given accommodation already rated by this User!")
	}

	err = service.rateAccommodationStore.Create(newRate)
	if err != nil {
		return err
	}

	return nil
}

func (service *RatingService) CreateUserRate(rating *domain.RatingUser) error {
	if rating.Rate < 1 || rating.Rate > 5 {
		return errors.New("The rating should be between 1 and 5!")
	}

	rating.Status = domain.Pending
	id, err := service.rateUserStore.Create(rating)
	rating.Id = id
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(rating)
	if err != nil {
		_ = service.rateUserStore.UpdateStatus(rating.Id, domain.Canceled)
		return err
	}
	return nil
}

func (service *RatingService) UpdateUserRating(host, user string, rate int32) error {
	return service.rateUserStore.Update(host, user, rate)
}

func (service *RatingService) DeleteUserRating(host, user string) error {
	return service.rateUserStore.Delete(host, user)
}

func (service *RatingService) RateAlreadyExists(host, user, id string) (bool, error) {
	ratingId, err := primitive.ObjectIDFromHex(id)
	_, err = service.rateUserStore.GetByHostAndUser(host, user, ratingId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service *RatingService) GetHostRates(host string) (*pb.HostRatesResponse, error) {
	ratings, err := service.rateUserStore.GetHostRates(host)
	if err != nil {
		return nil, err
	}

	return &pb.HostRatesResponse{
		Ratings: domain.UserRatingsToGrpcList(ratings),
	}, err
}

func (service *RatingService) UpdateStatus(id primitive.ObjectID, status domain.Status) error {
	return service.rateUserStore.UpdateStatus(id, status)
}

func (service *RatingService) UpdateAccommodationRate(updateRate domain.RatingAccommodation) error {
	_, err := service.rateAccommodationStore.GetByUserAndAccommodation(updateRate.User, updateRate.Accommodation)
	if err != nil {
		return errors.New("Given accommodation isn't rated by this User! CAN NOT UPDATE IT!")
	}

	if updateRate.Rate < 1 || updateRate.Rate > 5 {
		return errors.New("The rating should be between 1 and 5!")
	}

	err = service.rateAccommodationStore.Update(updateRate)
	if err != nil {
		return err
	}

	return nil
}

func (service *RatingService) DeleteAccommodationRate(userEmail, accommodationId string) error {
	rate, err := service.rateAccommodationStore.GetByUserAndAccommodation(userEmail, accommodationId)
	if err != nil {
		return errors.New("Given accommodation isn't rated by this User! CAN NOT DELETE IT!")
	}

	err = service.rateAccommodationStore.Delete(rate)
	if err != nil {
		return err
	}

	return nil
}

func (service *RatingService) GetAllAccommodationRates(accommodationId string) (*pb.AllAccommodationRatesResponse, error) {
	rates, err := service.rateAccommodationStore.GetAllByAccommodation(accommodationId)
	if err != nil {
		return nil, err
	}

	return &pb.AllAccommodationRatesResponse{
		Rates: domain.RateListToGrpcRateList(rates),
	}, nil
}

func (service *RatingService) GetAverageUserRate(host string) (*pb.AverageUserRatingResponse, error) {
	rates, err := service.rateUserStore.GetHostRates(host)
	if err != nil {
		return nil, err
	}
	return &pb.AverageUserRatingResponse{
		Average: averageUserRate(rates),
	}, nil
}

func (service *RatingService) GetAverageAccommodationRate(accommodationId string) (*pb.AverageRateAccommodationResponse, error) {
	rates, err := service.rateAccommodationStore.GetAllByAccommodation(accommodationId)
	if err != nil {
		return nil, err
	}

	return &pb.AverageRateAccommodationResponse{
		Average: getAverageRate(rates),
	}, nil
}

func getAverageRate(rates []domain.RatingAccommodation) float32 {
	var converted float32 = 0
	var counter float32 = 0

	for _, entity := range rates {
		converted += float32(entity.Rate)
		counter++
	}

	if counter == 0 {
		return 0
	}
	return converted / counter
}

func averageUserRate(rates []domain.RatingUser) float32 {
	var converted float32 = 0
	var counter float32 = 0

	for _, entity := range rates {
		converted += float32(entity.Rate)
		counter++
	}

	if counter == 0 {
		return 0
	}
	return converted / counter
}

func accommodationExists(accommodationId string) (*accommodation_service.SingleAccommodation, error) {
	accommodation := clients.NewAccommodationClient(fmt.Sprintf("%s:%s", config.NewConfig().AccommodationServiceHost, config.NewConfig().AccommodationServicePort))

	return accommodation.GetById(context.Background(), &accommodation_service.AccommodationIdRequest{
		Id: accommodationId,
	})
}
