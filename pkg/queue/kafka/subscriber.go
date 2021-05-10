package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type subscriber struct {
	reader *kafka.Reader
}

// SubscriberOption is a subscriber option.
type SubscriberOption func(*subscriber)

type SubscribeConfig struct {
	Topic string
	Group string
	Brokers []string
}

// NewSubscriber new a kafka subscriber.
func NewSubscriber(conf *SubscribeConfig, opts ...SubscriberOption) Subscriber {
	sub := &subscriber{}
	for _, o := range opts {
		o(sub)
	}
	sub.reader = kafka.NewReader(kafka.ReaderConfig{
		Topic:   conf.Topic,
		GroupID: conf.Group,
		Brokers: conf.Brokers,
	})
	return sub
}

func (s *subscriber) Subscribe(ctx context.Context, h Handler) error {
	for {
		msg, err := s.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}
		header := make(map[string]string, len(msg.Headers))
		for _, h := range msg.Headers {
			header[h.Key] = string(h.Value)
		}
		_ = h(context.Background(), Event{
			Key:        string(msg.Key),
			Payload:    msg.Value,
			Properties: header,
		})
		s.reader.CommitMessages(ctx, msg)
	}
}

func (s *subscriber) Close() error {
	return s.reader.Close()
}
