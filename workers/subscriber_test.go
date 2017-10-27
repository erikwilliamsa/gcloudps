package workers_test

import (
	"cloud.google.com/go/pubsub"
	"github.com/erikwilliamsa/gcloudps/workers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var callct = 0

type MockSubscriberClient struct {
	ReturnNil bool
}

var _ = Describe("CountMessageHandler #OnMessage", func() {
	Context("When OnMessage  is called", func() {
		It("Should consume until toggled off and return the number of messages consumed", func() {
			cmh := workers.NewCountMessageHandler()
			cmh.AutoAck = false
			cmh.OnMessage(&pubsub.Message{Data: []byte("message")})
			Expect(cmh.Count).To(Equal(1))

		})
	})

	Context("When OnMessage  is called 1000 times", func() {
		It("Should consume until toggled off and return the number of messages consumed", func() {
			cmh := workers.NewCountMessageHandler()
			cmh.AutoAck = false

			for i := 1; i <= 1000; i++ {
				cmh.OnMessage(&pubsub.Message{Data: []byte("message")})

			}
			Expect(cmh.Count).To(Equal(1000))

		})
	})
})
