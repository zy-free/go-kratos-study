package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"time"
)

type subscriber struct {
	reader sarama.ConsumerGroup
}

// SubscriberOption is a subscriber option.
type SubscriberOption func(*subscriber)

// NewSubscriber new a kafka subscriber.
func NewSubscriber(group string, brokers []string) (Subscriber,error) {
	sub := &subscriber{}
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0 // specify appropriate version
	config.Consumer.Return.Errors = true
	config.Consumer.MaxWaitTime = time.Millisecond * 250
	config.Consumer.MaxProcessingTime = 50 * time.Millisecond
	consumer, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return nil, err
	}
	sub.reader = consumer
	return sub,nil
}
//
//type onMessage func(message sarama.ConsumerMessage)
//
//type ConsumerGroupHandler struct {
//	onmessage onMessage
//}
//
//func (handler *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
//	return nil
//}
//
//func (handler *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
//
//func (handler *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	for message := range claim.Messages() {
//		if handler.onmessage != nil {
//			handler.onmessage(*message)
//			sess.MarkMessage(message, "")
//		}
//	}
//	return nil
//}

func (s *subscriber) Subscribe(ctx context.Context, h Handler) error {
	for {
		return nil
		//msg, err := s.reader.Consume()
		//if err != nil {
		//	return err
		//}
		//header := make(map[string]string, len(msg.Headers))
		//for _, h := range msg.Headers {
		//	header[h.Key] = string(h.Value)
		//}
		//_ = h(context.Background(), Event{
		//	Key:        string(msg.Key),
		//	Payload:    msg.Value,
		//	Properties: header,
		//})
		//s.reader.CommitMessages(ctx, msg)
	}
}

func (s *subscriber) Close() error {
	return s.reader.Close()
}
