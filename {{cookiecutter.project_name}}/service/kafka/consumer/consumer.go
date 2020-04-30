package consumer

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
}

func Setup(brokers []string, clientID string, config *sarama.Config) (*Consumer, error) {
	consumer := &Consumer{nil}
	consumerGroup, err := sarama.NewConsumerGroup(brokers, clientID, config)
	if err != nil {
		return nil, err
	}

	consumer.consumerGroup = consumerGroup
	return consumer, nil
}

func (consumer *Consumer) Consume(topics []string, handler sarama.ConsumerGroupHandler) {
	for {
		err := consumer.consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			log.Fatalf("Error from consumer: %s", err)
		}
	}
}

// Close is used to handle a gracefull shutdown of Kafka consumergroup
func (consumer *Consumer) Close() {
	consumer.Close()
}
