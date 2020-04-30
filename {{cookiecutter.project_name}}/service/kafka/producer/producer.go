package producer

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

func SetupAsyncProducer(brokers []string, config *sarama.Config) (sarama.AsyncProducer, error) {
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	go produce(producer)

	return producer, nil
}

func produce(producer sarama.AsyncProducer) {
	for {
		select {
		case result := <-producer.Successes():
			key, _ := result.Key.Encode()
			messagesProduced.WithLabelValues(result.Topic, string(key), string(success)).Inc()
			log.Debugf("%s message produced with key: %s", result.Topic, key)
		case err := <-producer.Errors():
			key, _ := err.Msg.Key.Encode()
			messagesProduced.WithLabelValues(err.Msg.Topic, string(key), string(failed)).Inc()
			log.Errorf("Failed to produce message: %s", err)
		}
	}
}
