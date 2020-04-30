package consumer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// State
type state string

const (
	// Success state
	success state = "success"
	// Failed state
	failed state = "failed"
)

var (
	messagesConsumed = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka",
		Name:      "messages_consumed",
		Help:      "The total number of Kafka messages consumed",
	}, []string{"key", "topic", "state"})
)
