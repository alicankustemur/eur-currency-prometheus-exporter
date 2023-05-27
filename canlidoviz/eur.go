package canlidoviz

import (
	"log"
	"net/http"
	"time"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
)

var currentCurrencyMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "canlidoviz_currency",
	Help: "Current Currency Price of canlidoviz.com",
}, []string{})

func setCurrentEur() {

	for {

		req, err := http.NewRequest("GET", "https://canlidoviz.com/doviz-kurlari/kapali-carsi", nil)
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

		doc.Find(".page").Each(func(i int, s *goquery.Selection) {
			eurText := strings.Replace(s.Find("tbody").Find("tr[data-code='EUR']").Find("td.text-primary").First().Text(), " ", "", -1)
			eurText = strings.Replace(eurText, "\n", "", -1)
			currentEur, err = strconv.ParseFloat(eurText, 64)
		})

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(5 * time.Second)
	}

}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
