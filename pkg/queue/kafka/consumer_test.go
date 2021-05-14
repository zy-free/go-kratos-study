package kafka

import (
	"context"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	r := NewConsumer(&ConsumerConfig{
		Topic:   testTopic,
		Group:   testGroup,
		Brokers: testBrokers,
	})
	time.AfterFunc(time.Second*3, func() {
		r.Close()
	})
	r.Consume(context.Background(), func(ctx context.Context, event Event) error {
		t.Logf("sub: key=%s value=%s header=%v", event.Key, event.Payload, event.Properties)
		return nil
	})
}
