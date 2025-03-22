package lprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var namespace string
var subsystem string

func Init(namespace_, subsystem_ string) {
	namespace = namespace_
	subsystem = subsystem_
}

func RegisterMetrics(reg prometheus.Registerer) {
	RegisterGoroutineCount(reg)
	RegisterMemStats(reg)
	RegisterHttpServerMetrics(reg)
}
