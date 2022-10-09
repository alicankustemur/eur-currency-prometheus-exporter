package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	eur = "eur"
)

func returnCurrentEur() float64 {
	res, err := http.Get("https://www.qnbfinansbank.enpara.com/hesaplar/doviz-ve-altin-kurlari")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var currentEur float64

	doc.Find(".enpara-gold-exchange-rates__table-item").Each(func(i int, s *goquery.Selection) {

		if s.Find("span").First().Text() == `EUR (€)` {

			eurText := strings.Split(s.Find("span").First().Next().Text(), " ")[0]
			eurText = strings.Replace(eurText, ",", ".", -1)
			currentEur, err = strconv.ParseFloat(eurText, 64)
		}
	})

	return currentEur
}

func main() {

	currentCurrencyMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "current_currency",
		Help: "Current Currency Price of Enpara",
	}, []string{})
	err := prometheus.Register(currentCurrencyMetric)

	if err != nil {
		return
	}

	currentCurrencyMetric.WithLabelValues().Add(returnCurrentEur())

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
