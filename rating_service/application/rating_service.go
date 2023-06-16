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
)

type RatingService struct {
	rateAccommodationStore domain.RatingAccommodationStore
}

func NewRatingService(accommodationStore domain.RatingAccommodationStore) *RatingService {
	return &RatingService{
		rateAccommodationStore: accommodationStore,
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

func accommodationExists(accommodationId string) (*accommodation_service.SingleAccommodation, error) {
	accommodation := clients.NewAccommodationClient(fmt.Sprintf("%s:%s", config.NewConfig().AccommodationServiceHost, config.NewConfig().AccommodationServicePort))

	return accommodation.GetById(context.Background(), &accommodation_service.AccommodationIdRequest{
		Id: accommodationId,
	})
}
