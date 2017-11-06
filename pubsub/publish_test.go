package pubsub_test

import (
	ps "cloud.google.com/go/pubsub"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockPublish struct{}

var pubCountResult = 0

func (p *MockPublish) Publish(message ps.Message) {
	pubCountResult++
}

var _ = Describe("#Publish", func() {

	Context("Implmentation of a Publish interface", func() {
		mp := MockPublish{}
		msg := ps.Message{Data: []byte("derp")}
		mp.Publish(msg)
		It("Should return a PublisherClient stuct", func() {
			Expect(pubCountResult).To(Equal(1))
		})
	})

})
