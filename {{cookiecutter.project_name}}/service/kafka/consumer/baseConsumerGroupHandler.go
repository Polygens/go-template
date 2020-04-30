package consumer

import (
	"github.com/Shopify/sarama"
)

// BaseConsumer Kafka consumer
type BaseConsumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *BaseConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *BaseConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// // ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// func (c *BaseConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
// 	for message := range claim.Messages() {
// 		//c.metrics.KafkaMessagesConsumed.WithLabelValues(message.Topic, string(message.Key), string(metric.Success)).Inc()
// 		log.Debugf("Message received: value = %s", string(message.Value))
// 		session.MarkMessage(message, "")
// 	}

// 	return nil
// }
