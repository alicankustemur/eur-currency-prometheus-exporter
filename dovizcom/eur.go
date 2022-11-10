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

func setCurrentEur() {

	for {

		req, err := http.NewRequest("GET", "https://kur.doviz.com/enpara", nil)
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

		doc.Find("#currencies").Each(func(i int, s *goquery.Selection) {
			eurText := strings.Split(s.Find("tbody").Find("tr").Next().Find("td.text-bold").Text(), ",")
			euroTextWithDot := eurText[0] + "." + eurText[1]
			currentEur, err = strconv.ParseFloat(euroTextWithDot, 64)
		})

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(2 * time.Second)
	}

}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
