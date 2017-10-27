package workers

import (
	"context"
	"fmt"
	"sync"

	ps "cloud.google.com/go/pubsub"
	"github.com/erikwilliamsa/gcloudps/pubsub"
)

var locker sync.Mutex

// CountMessageHandler used to count the number of messages consumed
type CountMessageHandler struct {
	// AutoAck  auto ack messages
	AutoAck bool
	// Count the number of messages consumed
	Count int
}

// NewCountMessageHandler Initilize the count message handler
// returns CountMessageHandler a new count message handler
func NewCountMessageHandler() *CountMessageHandler {
	return &CountMessageHandler{true, 0}
}

// OnMessage calls cmh to increment the number of messages.
func (cmh *CountMessageHandler) OnMessage(m *ps.Message) error {
	locker.Lock()
	if m != nil {
		cmh.Count++
		if cmh.AutoAck {
			m.Ack()
		}
	}
	fmt.Printf("\rConsumed:  %d", cmh.Count)
	locker.Unlock()
	return nil
}

// Subscribe consumes messages from the subscriptions
// sc contains the subscription, and the OnMessge Handler
func Subscribe(ctx context.Context, sc pubsub.SubscriberClient) {

	sc.Subscritpion.Receive(ctx, func(c context.Context, m *ps.Message) {
		sc.Handler.OnMessage(m)
	})
}
