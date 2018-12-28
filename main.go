package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
	"time"
)

// Config section for each site
type Config struct {
	Interval time.Duration
}

// Result contains result of the one/last check
type Result struct {
	StatusCode int
	Length     int
	Duration   time.Duration
	Hash       uint64
	Time       time.Time
}

// Site contains configuration and results
type Site struct {
	Config     Config
	LastResult Result
	Mutex      sync.Mutex
}

// Sitest has main parameters and attributes
type Sitest struct {
	ConfigFile    string
	ListenAddress string
	Sites         map[string]*Site
	Metrics       PromCollectors
	StartTime     time.Time
}

func main() {

	sitest := Sitest{StartTime: time.Now()}
	flag.StringVar(&sitest.ConfigFile, "config_file", "/etc/sitest/sitest.yaml", "path to config-file")
	flag.StringVar(&sitest.ListenAddress, "listen_addr", "0.0.0.0:8080", "listen address")
	flag.Parse()

	sitest.LoadConfig()
	sitest.Metrics.Register(prometheus.DefaultRegisterer)

	log.Printf("Start querying sites...")
	for site := range sitest.Sites {
		go sitest.Run(site)
	}

	http.Handle("/", sitest)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(sitest.ListenAddress, nil))

}
