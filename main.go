package main

import (
	"log"
	"net/http"

	"github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom"
	"github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom_kapalicarsi"
	"github.com/alicankustemur/eur-currency-prometheus-exporter/kuveytturk"
	"github.com/alicankustemur/eur-currency-prometheus-exporter/tcmb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	err := prometheus.Register(tcmb.CurrentEur())

	if err != nil {
		log.Fatal(err)
	}

	err = prometheus.Register(kuveytturk.CurrentEur())

	if err != nil {
		log.Fatal(err)
	}

	err = prometheus.Register(dovizcom.CurrentEur())

	if err != nil {
		log.Fatal(err)
	}
	err = prometheus.Register(dovizcom_kapalicarsi.CurrentEur())

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
