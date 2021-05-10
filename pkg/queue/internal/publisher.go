package kafka

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
)

// ErrEventFull is a message event chan full.
var ErrEventFull = errors.New("message event chan full")

var _ Publisher = (*publisher)(nil)

type publisher struct {
	brokers []string
	writer  sarama.SyncProducer
}

// NewPublisher new a kafka publisher.
func NewPublisher(brokers []string) (Publisher, error) {
	pub := &publisher{
		brokers: brokers,
	}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	writer, err := sarama.NewSyncProducer(brokers, config)
	pub.writer = writer
	return pub, err
}

func (p *publisher) Publish(ctx context.Context,event Event) error {
	//headers := make([]kafka.Header, 0, len(event.Properties))
	//for k, v := range event.Properties {
	//	headers = append(headers, kafka.Header{Key: k, Value: []byte(v)})
	//}
	_, _, err := p.writer.SendMessage(&sarama.ProducerMessage{
		Topic: event.Topic,
		Key:   sarama.StringEncoder(event.Key),
		Value: sarama.ByteEncoder(event.Payload),
	})
	return err
}

func (p *publisher) Close() error {
	return p.writer.Close()
}
