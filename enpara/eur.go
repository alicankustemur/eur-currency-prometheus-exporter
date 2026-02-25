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

var httpClient = &http.Client{Timeout: 30 * time.Second}

func setCurrentEur() {
	for {
		req, err := http.NewRequest("GET", "https://www.qnbfinansbank.enpara.com/hesaplar/doviz-ve-altin-kurlari", nil)
		if err != nil {
			log.Println("enpara: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("enpara: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println("enpara: error parsing html:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc.Find(".enpara-gold-exchange-rates__table-item").Each(func(i int, s *goquery.Selection) {
			if s.Find("span").First().Text() == `EUR (€)` {
				eurText := strings.Split(s.Find("span").First().Next().Text(), " ")[0]
				eurText = strings.Replace(eurText, ",", ".", -1)
				currentEur, err := strconv.ParseFloat(eurText, 64)
				if err == nil {
					currentCurrencyMetric.WithLabelValues().Set(currentEur)
				}
			}
		})

		time.Sleep(1 * time.Minute)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
