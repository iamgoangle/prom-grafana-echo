package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "hello_processed_ops_total",
		Help: "The total number of processed events",
	})

	http.Handle("/", promhttp.InstrumentHandlerCounter(
		promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hello_requests_total",
				Help: "Total number of hello-world requests by HTTP code.",
			},
			[]string{"code", "method"},
		),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			opsProcessed.Inc()
			time.Sleep(2 * time.Second)

			fmt.Fprint(w, "Hello, world!")
		}),
	))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9009", nil)
}