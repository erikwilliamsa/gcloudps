package workers

import (
	"fmt"
	"strconv"

	ps "cloud.google.com/go/pubsub"
	pubsub "github.com/erikwilliamsa/gcloudps/pubsub"
)

// PublisheWorker Sets the publisher to use
type PublisheWorker struct {
	Publish pubsub.Publish
}

// GenerateMessages generats a number of messages
func (pw *PublisheWorker) GenerateMessages(count int) error {
	for i := 0; i < count; i++ {

		msgdata := "Gnerated message " + strconv.Itoa(i+1) + " of " + strconv.Itoa(count)
		msg := &ps.Message{Data: []byte(msgdata)}

		fmt.Println(msgdata)
		pw.Publish.Publish(msg)

	}
	return nil
}
