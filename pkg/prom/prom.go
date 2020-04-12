package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	defaultMetricPath = "/metrics"
	defaultSubsystem  = "echo"
)

var (
	MetricTypeCounter   = "counter"
	MetricTypeGauge     = "gauge"
	MetricTypeHistogram = "histogram"
	MetricTypeSummary   = "summary"
)

var (
	stdCounter = &Metric{
		ID:          "stdCounter",
		Name:        "process_counter",
		Description: "How many counter processed",
		Type:        MetricTypeCounter,
	}

	stdGauge = &Metric{
		ID:          "stdGauge",
		Name:        "process_gauge",
		Description: "This is a gauge metric",
		Type:        MetricTypeGauge,
	}

	stdHistogram = &Metric{
		ID:          "stdHistogram",
		Name:        "process_histogram",
		Description: "The request histogram",
		Type:        MetricTypeHistogram,
	}

	stdSummary = &Metric{
		ID:          "stdSummary",
		Name:        "process_summary",
		Description: "The request summary",
		Type:        MetricTypeSummary,
	}
)

var standardMetrics = []*Metric{
	stdCounter,
	stdGauge,
	stdHistogram,
	stdSummary,
}

// Metric is a definition for the name, description, type, id, and
// prometheus.Collector type (i.e. CounterVec, Summary, etc) of each metric
type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

type PromClient struct {
	MetricPath  string
	MetricLists []*Metric

	Namespace string
	SubSystem string
}

func NewPrometheus(ns, subSys string, customMetricsList ...*Metric) *PromClient {
	var metricsList []*Metric

	if len(customMetricsList) > 1 {
		for _, metric := range customMetricsList {
			metricsList = append(metricsList, metric)
		}
	} else {
		metricsList = append(metricsList, customMetricsList[0])
	}

	for _, metric := range standardMetrics {
		metricsList = append(metricsList, metric)
	}

	return &PromClient{
		Namespace: ns,
		SubSystem: subSys,
		MetricLists: metricsList,
	}
}

func (p *PromClient) instanceMetric(m *Metric) prometheus.Collector {
	var metric prometheus.Collector

	switch m.Type {
	case MetricTypeCounter:
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: p.Namespace,
				Name:      m.Name,
				Help:      m.Description,
			})
	case MetricTypeGauge:
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: p.Namespace,
				Name:      m.Name,
				Help:      m.Description,
			})
	case MetricTypeHistogram:
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: p.Namespace,
				Name:      m.Name,
				Help:      m.Description,
			})
	case MetricTypeSummary:
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Namespace: p.Namespace,
				Name:      m.Name,
				Help:      m.Description,
			})
	}

	return metric
}

func (p *PromClient) applyMetrics() error {
	for _, m := range p.MetricLists {
		metricObject := p.instanceMetric(m)
		err := prometheus.Register(metricObject)
		if err != nil {
			return err
		}

		m.MetricCollector = metricObject
	}

	return nil
}
