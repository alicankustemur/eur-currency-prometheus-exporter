package dovizcom

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
	Name: "dovizcom_currency",
	Help: "Current Currency Price of Doviz.com",
}, []string{})

var httpClient = &http.Client{Timeout: 30 * time.Second}

func setCurrentEur() {
	for {
		req, err := http.NewRequest("GET", "https://kur.doviz.com/enpara", nil)
		if err != nil {
			log.Println("dovizcom: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("dovizcom: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println("dovizcom: error parsing html:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc.Find("#currencies").Each(func(i int, s *goquery.Selection) {
			eurText := strings.Split(s.Find("tbody").Find("tr").Next().Find("td.text-bold").Text(), ",")
			if len(eurText) >= 2 {
				euroTextWithDot := eurText[0] + "." + eurText[1]
				currentEur, err := strconv.ParseFloat(euroTextWithDot, 64)
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
