package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})
)

// Metrics for Prometheus
func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, _ := mux.CurrentRoute(r).GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, r.Method, r.Response.Status))
		defer timer.ObserveDuration()

		next.ServeHTTP(w, r)
	})
}
