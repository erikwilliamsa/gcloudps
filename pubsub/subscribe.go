package pubsub

type Subscriber interface{}
type subscriber struct {
}

func NewSubscriber() Subscriber {
	return subscriber{}
}
