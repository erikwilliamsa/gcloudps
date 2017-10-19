package workers_test

import (
	"fmt"
	"time"

	. "github.com/erikwilliamsa/gcloudps/workers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subscriber", func() {
	Context("When given a valid topic", func() {
		It("Should consume until toggled off and return the number of messages consumed", func() {

			toggle := make(chan struct{})
			c1 := make(chan int)

			go func() {
				c1 <- Consume(toggle)
			}()
			time.Sleep(time.Millisecond * 3)
			fmt.Println("Slept")
			close(toggle)
			count := <-c1

			Expect(count).ToNot(Equal(0))
		})
	})
})
