package main

import (
	"fmt"
	"net/http"
	"time"
	"flag"

	"github.com/prometheus/client_golang/prometheus"
   	"github.com/prometheus/client_golang/prometheus/promhttp"
    log "github.com/sirupsen/logrus"
)

type positiveDuration struct{ time.Duration }

var (
    printVersion = flag.Bool("version", false, "Print the version and exit")
    logFile     = flag.String("log.file", "/var/log/tinyproxy/tinyproxy.log", "Path to Tinyproxy log file")
    listenAddr  = flag.String("web.listen-address", ":8889", "Address to listen on for web interface and telemetry")

    connectionsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "tinyproxy_connections_open",
        Help: "Number of open connections",
    })

    requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "tinyproxy_requests_total",
        Help: "Total number of requests processed by Tinyproxy",
    })
)

func init() {
    prometheus.MustRegister(connectionsTotal)
    prometheus.MustRegister(requestsTotal)
}

func main() {
    flag.Parse()

    if *printVersion {
        fmt.Println("Tinyproxy Exporter Version 1.0")
        return
    }

    // Register metrics handler
    http.Handle("/metrics", promhttp.Handler())

    // Start server
    log.Infof("Starting HTTP server on %s", *listenAddr)
    if err := http.ListenAndServe(*listenAddr, nil); err != nil {
        log.Fatalf("Failed to start HTTP server: %s", err)
    }
}


