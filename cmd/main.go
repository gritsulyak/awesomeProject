package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	<-done
}
