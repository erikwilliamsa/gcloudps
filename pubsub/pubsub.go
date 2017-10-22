package pubsub

import "cloud.google.com/go/pubsub"

type SubcriberClient interface {
	Receive() pubsub.Message
}
