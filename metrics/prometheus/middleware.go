package prometheus

import (
	"net/http"
	"time"
	"github.com/spotxchange/hydra/health"
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
	if r.URL.Path != health.HealthMetricsPath &&
		r.URL.Path != health.HealthStatusPath &&
		r.URL.Path != health.PrometheusStatusPath {
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