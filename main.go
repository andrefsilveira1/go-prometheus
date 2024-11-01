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

func main() {
	fmt.Println("Starting Gauge metric")
	r := prometheus.NewRegistry()
	r.MustRegister(onlinesUsers)
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
