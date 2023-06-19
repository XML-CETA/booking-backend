package api

import (
	"context"
	"fmt"

	"booking-backend/accommodation_service/application"
	"booking-backend/accommodation_service/domain"
	"booking-backend/accommodation_service/startup/config"

	"booking-backend/common/clients"
	pb "booking-backend/common/proto/accommodation_service"
	"booking-backend/common/proto/auth_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllAccommodationRequest) (*pb.GetAllAccommodationResponse, error) {
	accommodations, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	return &pb.GetAllAccommodationResponse{
		Accommodations: accommodations,
	}, nil
}

func (handler *AccommodationHandler) GetById(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.SingleAccommodation, error) {
	accommodation, err := handler.service.GetById(request.Id)

	if err != nil {
		return nil, err
	}

	return accommodation, err
}

func (handler *AccommodationHandler) Create(ctx context.Context, request *pb.AccommodationCreateRequest) (*pb.Response, error) {
	user, err := Authorize(ctx, "HOST")
	if err != nil {
		return nil, err
	}

	newAccommodation := domain.MakeCreateAccommodation(request, user)

	err = handler.service.Create(newAccommodation)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Successfully Created!"),
	}, nil
}

func (handler *AccommodationHandler) Update(ctx context.Context, request *pb.SingleAccommodation) (*pb.Response, error) {
	acc, err := domain.MakeAccommodation(request)
	if err != nil {
		return nil, err
	}

	err = handler.service.Update(acc)
	if err != nil {
		return &pb.Response{
			Data: fmt.Sprint(err.Error()),
		}, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully updated"),
	}, nil
}

func (handler *AccommodationHandler) Delete(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.Response, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return &pb.Response{
			Data: fmt.Sprintf(err.Error()),
		}, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully deleted!"),
	}, nil
}

func (handler *AccommodationHandler) SearchAccommodations(ctx context.Context, request *pb.SearchAccommodationsRequest) (*pb.SearchAccommodationsResponse, error) {
	accommodations, err := handler.service.SearchAccommodations(request)
	if err != nil {
		return nil, err
	}

	return &pb.SearchAccommodationsResponse{
		Accommodations: accommodations,
	}, err
}

func (handler *AccommodationHandler) CreateAppointment(ctx context.Context, request *pb.SingleAppointment) (*pb.Response, error) {
	createAppointment, err := domain.MakeCreateAppointment(request)
	if err != nil {
		return nil, err
	}

	err = handler.service.AddAppointment(createAppointment)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Data: fmt.Sprintf("Succesfully created!"),
	}, nil
}

func (handler *AccommodationHandler) UpdateAppointment(ctx context.Context, request *pb.UpdateAppointmentRequest) (*pb.Response, error) {
	updateAppointment, err := domain.MakeUpdateAppointment(request)

	if err != nil {
		return nil, err
	}

	err = handler.service.UpdateAppointment(updateAppointment)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Data: fmt.Sprintf("Succesfully updated!"),
	}, nil
}

func (handler *AccommodationHandler) ValidateReservation(ctx context.Context, request *pb.ValidateReservationRequest) (*pb.ValidateReservationResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Accommodation)
	if err != nil {
		return nil, err
	}

	interval, err := domain.StringInvervalToDate(request.DateFrom, request.DateTo)
	if err != nil {
		return nil, err
	}

	fmt.Println("PROSO STRING TO DATE")
	fmt.Println(interval.DateFrom)
	fmt.Println(interval.DateTo)

	host, err := handler.service.ValidateReservation(id, interval)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateReservationResponse{
		Success: true,
		Host:    host,
	}, nil
}

func (handler *AccommodationHandler) IsAutomaticConfirmation(ctx context.Context, request *pb.AccommodationIdRequest) (*pb.IsAutomaticConfirmationResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	isAutomatic, err := handler.service.IsAutomaticConfirmation(id)
	if err != nil {
		return nil, err
	}
	return &pb.IsAutomaticConfirmationResponse{IsAutomaticConfirmation: isAutomatic}, nil
}

func Authorize(ctx context.Context, roleGuard string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

	return user.UserEmail, err
}
