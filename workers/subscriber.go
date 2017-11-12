package workers

import (
	"context"
	"fmt"
	"sync"

	ps "cloud.google.com/go/pubsub"
	f "github.com/erikwilliamsa/gcloudps/formatters"
	"github.com/erikwilliamsa/gcloudps/pubsub"
)

var locker sync.Mutex
var lastsize int

// CountMessageHandler used to count the number of messages consumed
type CountMessageHandler struct {
	// AutoAck  auto ack messages
	AutoAck bool
	// Count the number of messages consumed
	Count     int
	Preview   bool
	Formatter f.Formatter
}

// NewCountMessageHandler Initilize the count message handler
// returns CountMessageHandler a new count message handler
func NewCountMessageHandler() *CountMessageHandler {
	return &CountMessageHandler{true, 0, false, nil}
}

// OnMessage calls cmh to increment the number of messages.
func (cmh *CountMessageHandler) OnMessage(m *ps.Message) error {
	var msgdata string
	locker.Lock()
	if m != nil {
		if cmh.Preview {
			msgdata = cmh.Formatter.Format(string(m.Data))
		}
		cmh.Count++
		if cmh.AutoAck {
			m.Ack()
		}
	}

	if msgdata != "" {

		output := "\rConsumed: %d \n"
		fmt.Println("=============================================")
		fmt.Printf(output, cmh.Count)
		fmt.Println("\nMessage Preview:\n_____________________________________________")

		fmt.Printf("\n%s \n\n", msgdata)

		fmt.Println("=============================================")

	} else {
		fmt.Printf("\rConsumed: %d", cmh.Count)
	}
	locker.Unlock()
	return nil
}

// Subscribe consumes messages from the subscriptions
// sc contains the subscription, and the OnMessge Handler
func Subscribe(ctx context.Context, sc pubsub.SubscriberClient) {
	fmt.Println("=============================================")
	fmt.Println("\rConsumned: 0")

	sc.Subscritpion.Receive(ctx, func(c context.Context, m *ps.Message) {
		sc.Handler.OnMessage(m)
	})
}
