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

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "goapp_http_request_duration",
	Help: "Represents duration in seconds of all http requests",
}, []string{"handler"})

func main() {
	fmt.Println("Starting Gauge metric")
	r := prometheus.NewRegistry()
	r.MustRegister(onlinesUsers)
	r.MustRegister(httpRequests)
	r.MustRegister(httpDuration)

	home := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Handle"))
	})

	d := promhttp.InstrumentHandlerDuration(
		httpDuration.MustCurryWith(prometheus.Labels{"handler": "home"}),
		promhttp.InstrumentHandlerCounter(httpRequests, home),
	)

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.Handle("/metrics/counter", d)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
