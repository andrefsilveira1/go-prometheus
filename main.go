package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var onlinesUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "goapp_online_users",
	Help: "Users online",
	ConstLabels: map[string]string{
		"logged": "true",
	},
})

var httpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "goapp_http_requests_total",
	Help: "Count the amount of all HTTP request for goapp",
}, []string{})

func main() {
	fmt.Println("Starting Gauge metric")
	r := prometheus.NewRegistry()
	r.MustRegister(onlinesUsers)
	r.MustRegister(httpRequests)

	home := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Handle"))
	})
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.Handle("/metrics/counter", promhttp.InstrumentHandlerCounter(httpRequests, home))

	log.Fatal(http.ListenAndServe(":8181", nil))
}
