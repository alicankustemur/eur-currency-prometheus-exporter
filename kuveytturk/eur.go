package kuveytturk

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
	Name: "kuveytturk_currency",
	Help: "Current Currency Price of Kuveyt Turk",
}, []string{})

func setCurrentEur() {

	for {

		req, err := http.NewRequest("GET", "https://www.kuveytturk.com.tr/finans-portali/", nil)
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

		doc.Find(".col-sm-6").Each(func(i int, s *goquery.Selection) {

			if s.Find("h2").First().Text() == `EUR (Avrupa Para Birimi)` {
				eurText := s.Find("p").First().Text()
				eurText = strings.Replace(eurText, "\n", "", -1)
				eurText = strings.Replace(eurText, " ", "", -1)
				eurText = strings.Split(eurText, "Alış")[0]
				eurText = strings.Replace(eurText, ",", ".", -1)

				currentEur, err = strconv.ParseFloat(eurText, 64)
			}
		})

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(2 * time.Second)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
