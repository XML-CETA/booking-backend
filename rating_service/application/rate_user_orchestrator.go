package application

import (
	"booking-backend/common/messaging"
	events "booking-backend/common/saga/rate_user"
	"booking-backend/rating-service/domain"
)

type RateUserOrchestrator struct {
	commandPublisher messaging.PublisherModel
	replySubscriber  messaging.SubscriberModel
}

func NewRateUserOrchestrator(publisher messaging.PublisherModel, subscriber messaging.SubscriberModel) (*RateUserOrchestrator, error) {
	o := &RateUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *RateUserOrchestrator) Start(rating *domain.RatingUser) error {
	event := &events.RateUserCommand{
		Type: events.CheckUserVisit,
		Rating: events.UserRating{
			Id:        rating.Id.Hex(),
			RatedUser: rating.RatedUser,
			RatedBy:   rating.RatedBy,
			Rate:      rating.Rate,
		},
	}
	return o.commandPublisher.Publish(event)
}

func (o *RateUserOrchestrator) handle(reply *events.RateUserReply) {
	command := events.RateUserCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *RateUserOrchestrator) nextCommandType(reply events.RateUserReplyType) events.RateUserCommandType {
	switch reply {
	case events.UserRateDoesntExists:
		return events.ApproveRating
	case events.UserRateExists:
		return events.CancelRating
	case events.UserVisitedHost:
		return events.CheckUserRateExists
	case events.UserNotVisitedHost:
		return events.CancelRating
	default:
		return events.UnknownCommand
	}
}
