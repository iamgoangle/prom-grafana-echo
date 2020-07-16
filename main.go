package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "golf_http_metric_handler_requests_total",
			Help: "Total number of scrapes by HTTP status code.",
		},
		[]string{"code", "method", "path"},
	)

	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		// Namespace: config.Namespace,
		// Subsystem: config.Subsystem,
		Name: "golf_http_metric_requests_duration",
		Help: "Spend time by processing a route",
		Buckets: []float64{
			0.0005,
			0.001, // 1ms
			0.002,
			0.005,
			0.01, // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
			10.0, // 10s
			15.0,
			20.0,
			30.0,
		},
	}, []string{"method", "handler"})
)

func main() {
	runHTTPRequestCount()
	runHTTPRequestDuration()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9009", nil)
}

func runHTTPRequestDuration() {
	go func() {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(http.MethodGet, "/v1/golf/test"))
		defer func(){
			timer.ObserveDuration()
		}()

		time.Sleep(60 * time.Second)
	}()
}

func runHTTPRequestCount() {
	go func() {
		for {
			httpRequestsCount.WithLabelValues("200", http.MethodGet, "/v1/golf/test").Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
