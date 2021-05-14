package service

import (
	"context"

	"go-kartos-study/app/job/member/conf"
	"go-kartos-study/pkg/log"
	"go-kartos-study/pkg/queue/kafka"
)

// Service .
type Service struct {
	c        *conf.Config
	consumer kafka.Consumer
}

// New init service.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:        c,
		consumer: kafka.NewConsumer(c.KafkaConsumer),
	}

	go s.memberConsume(context.Background())

	return s
}

// Close service
func (s *Service) Close() {
}

// Ping service
func (s *Service) Ping(c context.Context) (err error) {
	return
}

func (s *Service) memberConsume(ctx context.Context) {
	defer s.consumer.Close()
	s.consumer.Consume(ctx, func(ctx context.Context, event kafka.Event) error {
		log.Infoc(ctx,"sub: key=%s value=%s header=%v", event.Key, event.Payload, event.Properties)
		return nil
	})
}

