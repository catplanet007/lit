package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	litprom "github.com/catplanet007/lit/lprometheus"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	randomNum := rand.Intn(100)
	time.Sleep(time.Duration(randomNum) * time.Millisecond)
	randomNum = rand.Intn(100)
	if randomNum < 30 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func main() {
	reg := prometheus.NewRegistry()
	litprom.Init("xx", "yy")
	litprom.RegisterMetrics(reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.Handle("/hello", litprom.HttpMetricsMiddleware(http.HandlerFunc(HelloHandler)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
