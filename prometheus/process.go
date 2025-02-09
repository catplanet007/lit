package prometheus

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func RegisterGoroutineCount(reg prometheus.Registerer) {
	goroutineCount := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "goroutine_cnt",
		Help:      "goroutine count",
	})
	go func() {
		cnt := runtime.NumGoroutine()
		goroutineCount.Set(float64(cnt))
		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()
		for range ticker.C {
			cnt := runtime.NumGoroutine()
			goroutineCount.Set(float64(cnt))
		}
	}()
	reg.MustRegister(goroutineCount)
}

func RegisterMemStats(reg prometheus.Registerer) {
	heapAlloc := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_heap_alloc_bytes",
		Help: "Number of bytes allocated on the heap",
	})
	stackInUse := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_stack_inuse_bytes",
		Help: "Number of bytes in use by stack allocations",
	})
	gcCount := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "go_gc_count_total",
		Help: "Total number of garbage collections",
	})

	reg.MustRegister(heapAlloc)
	reg.MustRegister(stackInUse)
	reg.MustRegister(gcCount)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		var lastGC uint32
		for range ticker.C {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			heapAlloc.Set(float64(memStats.HeapAlloc))
			stackInUse.Set(float64(memStats.StackInuse))
			if memStats.NumGC > lastGC {
				gcCount.Add(float64(memStats.NumGC - lastGC))
				lastGC = memStats.NumGC
			}
		}
	}()
}
