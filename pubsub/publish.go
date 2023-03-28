package pubsub

import (
	"context"
	"fmt"

	ps "cloud.google.com/go/pubsub"
)

// Publish interface intended to allow mocking of google Topic.Publish()
type Publish interface {
	Publish(message *ps.Message)
}

// PublishClient settings for the Publisher
type PublishClient struct {
	Context context.Context
	Topic   *ps.Topic
	Batch   bool
}

// Publish Publishes a message to the provided topic.
func (pc *PublishClient) Publish(message *ps.Message) {
	r := pc.Topic.Publish(pc.Context, message)
	if !pc.Batch {
		_, err := r.Get(pc.Context)
		if err != nil {
			fmt.Printf("Error while publishing! %s ", err)
		}
	}

}
