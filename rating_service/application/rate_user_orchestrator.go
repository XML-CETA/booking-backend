package application

import (
	nats "booking-backend/common/messaging"
	events "booking-backend/common/saga/rate_user"
	"booking-backend/rating-service/domain"
)

type RateUserOrchestrator struct {
	commandPublisher nats.Publisher
	replySubscriber  nats.Subscriber
}

func NewRateUserOrchestrator(publisher nats.Publisher, subscriber nats.Subscriber) (*RateUserOrchestrator, error) {
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
		Type: events.ApproveRating,
		Rating: events.UserRating{
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
