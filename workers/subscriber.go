package workers

import (
	ps "github.com/erikwilliamsa/gcloudps/pubsub"
)

// ConsumeCount until toggled off.
// It returns the number of messages consumed.
func ConsumeCount(toggle <-chan struct{}, s ps.SubcriberClient) int {
	i := 0
	for {
		select {
		case <-toggle:
			return i
		default:
			s.Receive()
			i++
		}

	}
}
