package lprometheus

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var requestDuration *prometheus.HistogramVec
var httpRequests *prometheus.CounterVec

func RegisterHttpServerMetrics(reg prometheus.Registerer) {
	requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "http_request_duration_seconds",
		Help:      "Histogram of http request processing time in millseconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"method", "path"})

	httpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "http_requests_total",
		Help:      "Total number of http requests, labeled by method, path and status code",
	}, []string{"method", "path", "code"})

	reg.MustRegister(requestDuration)
	reg.MustRegister(httpRequests)
}

type wrapedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ww *wrapedResponseWriter) WriteHeader(code int) {
	ww.statusCode = code
	ww.ResponseWriter.WriteHeader(code)
}

func HttpMetricsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &wrapedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(ww, r)
		duration := time.Since(start)
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(float64(duration.Seconds()))
		statusCodeStr := strconv.Itoa(ww.statusCode)
		httpRequests.WithLabelValues(r.Method, r.URL.Path, statusCodeStr).Inc()
	})
}
