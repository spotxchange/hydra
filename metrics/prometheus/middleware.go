package prometheus

import (
	"net/http"
	"time"
)

type MetricsManager struct {
	prometheusMetrics *Metrics
}

func NewMetricsManager(version, hash, buildTime string) *MetricsManager {
	return &MetricsManager{
		prometheusMetrics: NewMetrics(version, hash, buildTime),
	}
}

func (pmm *MetricsManager) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	if r.URL.Path != "/health/status" &&
		r.URL.Path != "/health/metrics" &&
		r.URL.Path != "/health/prometheus" {
		// Request counter metric
		pmm.prometheusMetrics.Counter.WithLabelValues(r.URL.Path).Inc()
		// Response time metric
		pmm.prometheusMetrics.ResponseTime.WithLabelValues(r.URL.Path).Observe(time.Since(start).Seconds())
	}
}

func (pmm *MetricsManager) Reset() {
	pmm.prometheusMetrics.Counter.Reset()
	pmm.prometheusMetrics.ResponseTime.Reset()
}