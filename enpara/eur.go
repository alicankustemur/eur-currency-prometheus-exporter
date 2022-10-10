package enpara

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
)

var currentCurrencyMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "enpara_currency",
	Help: "Current Currency Price of Enpara",
}, []string{})

func setCurrentEur() {

	for {

		req, err := http.NewRequest("GET", "https://www.qnbfinansbank.enpara.com/hesaplar/doviz-ve-altin-kurlari", nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		var currentEur float64

		doc.Find(".enpara-gold-exchange-rates__table-item").Each(func(i int, s *goquery.Selection) {

			if s.Find("span").First().Text() == `EUR (€)` {

				eurText := strings.Split(s.Find("span").First().Next().Text(), " ")[0]
				eurText = strings.Replace(eurText, ",", ".", -1)
				currentEur, err = strconv.ParseFloat(eurText, 64)
			}
		})

		time.Sleep(2 * time.Second)

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
	}

}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
