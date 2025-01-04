package monitoing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// common method get gauge metrics alongwith labels
func getGaugeMetrics(name string, help string) prometheus.Gauge {
	return prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		})
}

var latencyMetrics = getGaugeMetrics("go_network_latency", "latency Metrics From custom GO app")

func StartMetricsCollector(metricsCollectorPort string) {
	metricsCollectorPort = ":" + metricsCollectorPort
	log.Printf("Starting metric collector @ %s", metricsCollectorPort)
	prometheus.MustRegister(latencyMetrics)

	http.Handle("/metrics", newHttpHandler(promhttp.Handler()))
	err := http.ListenAndServe(metricsCollectorPort, nil)
	if err != nil {
		fmt.Printf("Failed to start metric collector @ %s", metricsCollectorPort)
	}
	fmt.Println("Started Prometheus server at", metricsCollectorPort)
}

// send http request
func newHttpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// start := time.Now()
		status := http.StatusOK

		// defer func() {
		// 	histogram.WithLabelValues(fmt.Sprintf("%d", status)).Observe(time.Since(start).Seconds())
		// }()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
		status = http.StatusBadRequest

		w.WriteHeader(status)
	})
}
