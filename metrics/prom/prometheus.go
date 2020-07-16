package prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricTypeCounter   = "counter"
	MetricTypeGauge     = "gauge"
	MetricTypeHistogram = "histogram"
	MetricTypeSummary   = "summary"
)

type Metric struct {
	Name string
	Help string
	Type string
}

type Prom struct {
	MetricList []*Metric

	Counter      prometheus.Counter
	CounterVec   *prometheus.CounterVec
	CounterFunc  prometheus.CounterFunc
	HistogramVec *prometheus.HistogramVec
}

func NewProm() *Prom {
	// promauto.NewCounter(prometheus.CounterOpts{
	// 	Name: "myapp_processed_ops_total",
	// 	Help: "The total number of processed events",
	// })
	return &Prom{}
}

func (p *Prom) registerStandardMetric() {
	p.Counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	
	p.HistogramVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_handlers_duration_seconds",
		Help: "Handlers request duration in seconds",
	}, []string{"path"})
}
