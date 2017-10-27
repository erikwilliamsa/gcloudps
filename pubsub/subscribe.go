package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

/*
	MessageHandler interface
	OnMessage function is what a subscriber will call. Any work done on a message
	should be done in the OnMessage function
*/
type MessageHandler interface {
	OnMessage(*pubsub.Message) error
}

/*
	SubscriberClient Holdes the methods that will be used in the workers
*/
type SubscriberClient struct {
	Subscritpion *pubsub.Subscription
	Context      context.Context
	Handler      MessageHandler
}

/*
	NewSubscriberClient initilizes a new subscriber to be uused in a worker.
*/
func NewSubscriberClient(ctx context.Context, s *pubsub.Subscription, mh MessageHandler) SubscriberClient {
	return SubscriberClient{s, ctx, mh}
}
