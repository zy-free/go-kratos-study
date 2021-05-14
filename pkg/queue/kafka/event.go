package kafka

import "context"

type Event struct {
	// Key sets the key of the message for routing policy
	Key string
	// Payload for the message
	Payload []byte
	// Properties attach application defined properties on the message
	Properties map[string]string
}

// Handler is a callback function that processes messages delivered
// to asynchronous subscribers.
type Handler func(context.Context, Event) error

// Publisher is absctraction for sending messages
// to queue.
type Publisher interface {
	Publish(ctx context.Context, event Event) error
	Close() error
}

// Consumer is an absctraction for receiving messages
// from queue.
type Consumer interface {
	Consume(ctx context.Context, h Handler) error
	Close() error
}
