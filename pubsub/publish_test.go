package pubsub_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockPublish struct{}

var pubCountResult = 0

func (p *MockPublish) Publish(message string) {
	pubCountResult++
}

var _ = Describe("#Publish", func() {

	Context("Implmentation of a Publish interface", func() {
		mp := MockPublish{}
		mp.Publish("derp")
		It("Should return a PublisherClient stuct", func() {
			Expect(pubCountResult).To(Equal(1))
		})
	})
})
