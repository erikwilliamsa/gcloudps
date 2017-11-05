package workers_test

import (
	ps "cloud.google.com/go/pubsub"
	workers "github.com/erikwilliamsa/gcloudps/workers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockpublish struct {
	total int
}

func (mp *mockpublish) Publish(message *ps.Message) {

	mp.total++
}

var _ = Describe("#GenerateMessage", func() {
	Context("When a positive number is given ", func() {
		msgcount := 4
		mp := mockpublish{0}
		pw := workers.PublisheWorker{&mp}

		It("Should publish the amount requested ", func() {
			pw.GenerateMessages(msgcount)
			Expect(mp.total).To(Equal(msgcount))
		})
	})
})
