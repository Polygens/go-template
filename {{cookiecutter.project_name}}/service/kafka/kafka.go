package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
	"github.com/polygens/{{cookiecutter.project_name}}/service/kafka/producer"
	{% endif -%}
	{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
	"github.com/polygens/{{cookiecutter.project_name}}/service/kafka/consumer"
	{% endif -%}
)

type Kafka struct {
	{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
	Consumer *consumer.Consumer
	{% endif -%}
	{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
	Producer sarama.AsyncProducer
	{%- endif %}
}

func Setup(cfg *config.Kafka) (*Kafka, error){
	kafka := &Kafka{ {% if cookiecutter.kafka|lower in ["consumer", "both"] -%}nil, {%- endif %}{% if cookiecutter.kafka|lower in ["producer", "both"] -%}nil, {%- endif %}}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.ClientID = cfg.ClientID

	var err error
	config.Version, err = sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		log.Fatalf("Invalid kafka version used: %s", err)
	}
	
	{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
	kafka.Producer, err = producer.SetupAsyncProducer(cfg.Brokers, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create producer: %w", err)
	}
	{%- endif %}
	
	{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
	kafka.Consumer, err = consumer.Setup(cfg.Brokers, cfg.ClientID, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create consumer: %w", err)
	}
	{%- endif %}
	
	return kafka, nil
}

// Close is used to handle a gracefull shutdown of kafka
func (kafka *Kafka) Close() {
	{% if cookiecutter.kafka|lower in ["consumer", "both"] -%}
	kafka.Consumer.Close()
	{% endif -%}
	{% if cookiecutter.kafka|lower in ["producer", "both"] -%}
	kafka.Producer.AsyncClose()
	{%- endif %}
}
