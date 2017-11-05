package pubsub

// Publish interface intended to allow mocking of google Topic.Publish()
type Publish interface {
	Publish(message string)
}
