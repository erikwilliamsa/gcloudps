package workers

import (
	"fmt"

	ps "cloud.google.com/go/pubsub"
	pubsub "github.com/erikwilliamsa/gcloudps/pubsub"
)

type PublisheWorker struct {
	Publish pubsub.Publish
}

func (pw *PublisheWorker) GenerateMessages(count int) error {
	for i := 0; i < count; i++ {
		msgdata := "Gnerated message " + string(i+1) + " of " + string(count)
		msg := &ps.Message{Data: []byte(msgdata)}

		fmt.Println(msgdata)
		pw.Publish.Publish(msg)

	}
	return nil
}
