package kafka

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

// ErrEventFull is a message event chan full.
var ErrEventFull = errors.New("message event chan full")

var _ Publisher = (*publisher)(nil)

const (
	// RequireNone the producer won’t even wait for a response from the broker.
	RequireNone kafka.RequiredAcks = kafka.RequireNone
	// RequireOne the producer will consider the write successful when the leader receives the record.
	RequireOne kafka.RequiredAcks = kafka.RequireOne
	// RequireAll the producer will consider the write successful when all of the in-sync replicas receive the record.
	RequireAll kafka.RequiredAcks = kafka.RequireAll
)

// PublisherOption is a publisher options.
type PublisherOption func(*publisher)

// ReadTimeout with read timeout option.
func ReadTimeout(d time.Duration) PublisherOption {
	return func(o *publisher) {
		o.readTimeout = d
	}
}

// WriteTimeout with write timeout option.
func WriteTimeout(d time.Duration) PublisherOption {
	return func(o *publisher) {
		o.writeTimeout = d
	}
}

// RequiredAcks with required acks option.
func RequiredAcks(acks kafka.RequiredAcks) PublisherOption {
	return func(o *publisher) {
		o.requiredAcks = acks
	}
}

type pubEvent struct {
	ctx      context.Context
	event    Event
	callback func(err error)
}

type publisher struct {
	brokers      []string
	readTimeout  time.Duration
	writeTimeout time.Duration
	requiredAcks kafka.RequiredAcks
	writer       *kafka.Writer
}

type PublishConfig struct {
	Topic string
	Brokers []string
}

// NewPublisher new a kafka publisher.
func NewPublisher(conf *PublishConfig, opts ...PublisherOption) Publisher {
	pub := &publisher{
		brokers:      conf.Brokers,
		readTimeout:  500 * time.Millisecond,
		writeTimeout: 500 * time.Millisecond,
		requiredAcks: RequireOne,
	}
	for _, o := range opts {
		o(pub)
	}
	pub.writer = &kafka.Writer{
		Addr:         kafka.TCP(conf.Brokers...),
		Topic:        conf.Topic,
		Balancer:     &kafka.Hash{},
		RequiredAcks: pub.requiredAcks,
		ReadTimeout:  pub.readTimeout,
		WriteTimeout: pub.writeTimeout,
	}
	return pub
}

func (p *publisher) Publish(ctx context.Context, event Event) error {
	headers := make([]kafka.Header, 0, len(event.Properties))
	for k, v := range event.Properties {
		headers = append(headers, kafka.Header{Key: k, Value: []byte(v)})
	}
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:     []byte(event.Key),
		Value:   event.Payload,
		Headers: headers,
	})
}

func (p *publisher) Close() error {
	return p.writer.Close()
}
