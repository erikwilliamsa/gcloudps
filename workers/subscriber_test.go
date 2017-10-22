package workers_test

import (
	"time"

	"cloud.google.com/go/pubsub"
	. "github.com/erikwilliamsa/gcloudps/workers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var callct = 0

type MockSubscriberClient struct{}

func (msc *MockSubscriberClient) Receive() pubsub.Message {
	callct++
	var b []byte
	return pubsub.Message{Data: b}
}

var _ = Describe("Subscriber #ConsumeCount", func() {
	Context("When consume count is called", func() {
		It("Should consume until toggled off and return the number of messages consumed", func() {
			mc := &MockSubscriberClient{}
			toggle := make(chan struct{})
			c1 := make(chan int)

			go func() {
				c1 <- ConsumeCount(toggle, mc)
			}()
			time.Sleep(time.Millisecond * 1) // Make sure it can go at least 1 iteration.
			close(toggle)
			count := <-c1

			Expect(count).ToNot(Equal(0))
			Expect(count).To(Equal(callct))
		})
	})
})
