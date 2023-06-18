package api

import (
	"booking-backend/common/messaging"
	events "booking-backend/common/saga/rate_user"
	"booking-backend/rating-service/application"
	"booking-backend/rating-service/domain"
)

type RateUserCommandHandler struct {
	ratingService     *application.RatingService
	replyPublisher    messaging.PublisherModel
	commandSubscriber messaging.SubscriberModel
}

func NewRateUserCommandHandler(ratingService *application.RatingService, publisher messaging.PublisherModel, subscriber messaging.SubscriberModel) (*RateUserCommandHandler, error) {
	o := &RateUserCommandHandler{
		ratingService:     ratingService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
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
	case events.CheckUserRateExists:
		rateExists, _ := handler.ratingService.RateAlreadyExists(command.Rating.RatedUser, command.Rating.RatedBy, command.Rating.Id)
		if rateExists {
			reply.Type = events.UserRateExists
			break
		}
		reply.Type = events.UserRateDoesntExists
	case events.ApproveRating:
		err := handler.ratingService.UpdateStatus(command.Rating.RatedUser, command.Rating.RatedBy, domain.Approved)

		if err != nil {
			return
		}
		reply.Type = events.RatingApproved
	case events.CancelRating:
		err := handler.ratingService.UpdateStatus(command.Rating.RatedUser, command.Rating.RatedBy, domain.Canceled)
		if err != nil {
			return
		}
		reply.Type = events.RatingCanceled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
