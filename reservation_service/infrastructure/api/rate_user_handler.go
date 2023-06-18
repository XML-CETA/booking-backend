package api

import (
	"booking-backend/common/messaging"
	events "booking-backend/common/saga/rate_user"
	"booking-backend/reservation-service/application"
)

type RateUserCommandHandler struct {
	reservationService *application.ReservationService
	replyPublisher     messaging.PublisherModel
	commandSubscriber  messaging.SubscriberModel
}

func NewRateUserCommandHandler(reservationService *application.ReservationService, publisher messaging.PublisherModel, subscriber messaging.SubscriberModel) (*RateUserCommandHandler, error) {
	o := &RateUserCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *RateUserCommandHandler) handle(command *events.RateUserCommand) {
	reply := events.RateUserReply{Rating: command.Rating}

	switch command.Type {
	case events.CheckUserVisit:
		visited, _ := handler.reservationService.CheckIfUserVisitedHost(command.Rating.RatedBy, command.Rating.RatedUser)
		if visited {
			reply.Type = events.UserVisitedHost
			break
		}
		reply.Type = events.UserNotVisitedHost
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
