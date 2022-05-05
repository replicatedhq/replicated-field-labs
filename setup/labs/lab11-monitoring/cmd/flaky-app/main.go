package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)


var (
	cpuTempGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "Current temperature.",
	})
	cpuTemp = 79.8
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTempGauge)
}

func main() {
	cpuTempGauge.Set(cpuTemp)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/temp", handleTemp)
	http.HandleFunc("/healthz", handleHealthz)
	http.HandleFunc("/favicon.ico", handleHealthz)
	http.Handle("/metrics", promhttp.Handler())

	log.Print("starting server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
