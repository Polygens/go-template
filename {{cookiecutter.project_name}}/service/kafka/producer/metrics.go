package producer

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
	messagesProduced = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka",
		Name:      "messages_produced",
		Help:      "The total number of Kafka messages produced",
	}, []string{"topic", "key", "state"})
)
