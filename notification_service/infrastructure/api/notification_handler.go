package api

import (
	"booking-backend/common/proto/notification_service"
	"booking-backend/notification_service/application"
	"context"
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

func (h NotificationHandler) Hello(ctx context.Context, request *notification_service.GetUserNotificationsRequest) (*notification_service.GetUserNotificationsResponse, error) {
  return &notification_service.GetUserNotificationsResponse{
  }, nil
}
