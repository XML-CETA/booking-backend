package messaging

type PublisherModel interface {
	Publish(message interface{}) error
}

type SubscriberModel interface {
	Subscribe(function interface{}) error
}

