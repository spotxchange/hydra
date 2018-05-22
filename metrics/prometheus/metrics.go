package prometheus

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct{
	Counter      *prom.CounterVec
	ResponseTime *prom.HistogramVec
}

func NewMetrics(version, hash, buildTime string) *Metrics {
	pm := &Metrics{
		Counter: prom.NewCounterVec(
			prom.CounterOpts{
				Name:        "hydra_requests_total",
				Help:        "Description",
				ConstLabels: map[string]string{
					"version":   version,
					"hash":      hash,
					"buildTime": buildTime,
				},
			},
			[]string{"endpoint"},
		),
		ResponseTime: prom.NewHistogramVec(
			prom.HistogramOpts{
				Name:        "hydra_response_time_seconds",
				Help:        "Description",
				ConstLabels: map[string]string{
					"version":   version,
					"hash":      hash,
					"buildTime": buildTime,
				},
			},
			[]string{"endpoint"},
		),
	}
	prom.Register(pm.Counter)
	prom.Register(pm.ResponseTime)
	return pm
}