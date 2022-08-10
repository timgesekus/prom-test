package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			queueLength.Set(rand.Float64()*20 + 79)
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cicis_processed_messages",
		Help: "Total number of proccesed_messages",
	})
	queueLength = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cicis_message_queue_size",
		Help: "Size of message queue in percent",
	})
)

func main() {
	recordMetrics()
	rand.Seed(time.Now().UnixNano())
	prometheus.MustRegister(opsProcessed)
	prometheus.MustRegister(queueLength)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
