package api

import (
	"booking-backend/common/clients"
	"booking-backend/common/proto/auth_service"
	"booking-backend/common/proto/notification_service"
	"booking-backend/notification_service/application"
	"booking-backend/notification_service/startup/config"
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type NotificationHandler struct {
  notification_service.UnimplementedNotificationServiceServer
	service *application.NotificationService
}

func NewNotificationHandler(service *application.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (h NotificationHandler) GetUserNotifications(ctx context.Context, request *notification_service.GetUserNotificationsRequest) (*notification_service.GetUserNotificationsResponse, error) {
	user, err := Authorize(ctx, "HOST")
	if err != nil {
		return nil, err
	}

  settings, err := h.service.GetUserSettings(user)
	if err != nil {
		return nil, err
	}


  return &notification_service.GetUserNotificationsResponse{
      User: settings.User,
      Role: settings.Role,
      ReservationRequest: settings.ReservationRequest,
      ReservationCancel: settings.ReservationCancel,
      PersonalRating: settings.PersonalRating,
      AccommodationRating: settings.AccommodationRating,
      ProminentStatusChange: settings.ProminentStatusChange,
      ReservationResponse: settings.ReservationResponse,
  }, nil
}

func (h NotificationHandler) NewUserSettings(ctx context.Context, request *notification_service.NewUserSettingsRequest) (*notification_service.NewUserSettingsResponse, error) {
  err := h.service.NewUserSettings(request.Host, request.Role)
  if err != nil {
    return nil, err
  }

  return &notification_service.NewUserSettingsResponse{
  }, nil
}

func Authorize(ctx context.Context, roleGuard string) (string, error) {
	auth := clients.NewAuthClient(fmt.Sprintf("%s:%s", config.NewConfig().AuthServiceHost, config.NewConfig().AuthServicePort))
	md, _ := metadata.FromIncomingContext(ctx)
	user, err := auth.Authorize(metadata.NewOutgoingContext(ctx, md), &auth_service.AuthorizeRequest{RoleGuard: roleGuard})

  if err != nil {
    return "", err
  }

	return user.UserEmail, nil
}
