package canlidoviz

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
	Name: "canlidoviz_currency",
	Help: "Current Currency Price of canlidoviz.com",
}, []string{})

var httpClient = &http.Client{Timeout: 30 * time.Second}

func setCurrentEur() {
	for {
		req, err := http.NewRequest("GET", "https://canlidoviz.com/doviz-kurlari/kapali-carsi", nil)
		if err != nil {
			log.Println("canlidoviz: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("canlidoviz: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println("canlidoviz: error parsing html:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc.Find(".page").Each(func(i int, s *goquery.Selection) {
			eurText := strings.Replace(s.Find("tbody").Find("tr[data-code='EUR']").Find("td.text-primary").First().Text(), " ", "", -1)
			eurText = strings.Replace(eurText, "\n", "", -1)
			eurText = strings.Replace(eurText, ",", ".", -1)
			currentEur, err := strconv.ParseFloat(eurText, 64)
			if err == nil {
				currentCurrencyMetric.WithLabelValues().Set(currentEur)
			}
		})

		time.Sleep(1 * time.Minute)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
