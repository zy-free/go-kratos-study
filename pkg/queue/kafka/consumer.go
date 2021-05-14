package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type consumer struct {
	reader *kafka.Reader
}

// ConsumerOption is a consumer option.
type ConsumerOption func(*consumer)

type ConsumerConfig struct {
	Topic string
	Group string
	Brokers []string
}

// NewConsumer new a kafka consumer.
func NewConsumer(conf *ConsumerConfig, opts ...ConsumerOption) Consumer {
	sub := &consumer{}
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

func (s *consumer) Consume(ctx context.Context, h Handler) error {
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

func (s *consumer) Close() error {
	return s.reader.Close()
}
