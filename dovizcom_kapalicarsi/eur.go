package dovizcom_kapalicarsi

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
	Name: "dovizcom_kapalicarsi_currency",
	Help: "Current Currency Price of Doviz.com",
}, []string{})

func setCurrentEur() {

	for {

		req, err := http.NewRequest("GET", "https://kur.doviz.com/kapalicarsi", nil)
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

		doc.Find("div.table,div.sortable").Each(func(i int, div *goquery.Selection) {

			if i == 0 {
				div.Find("tr").Each(func(j int, tr *goquery.Selection) {
					if j == 2 {
						tr.Find("td").Each(func(k int, td *goquery.Selection) {
							if k == 1 {
								eurText := strings.Split(td.Text(), ",")
								euroTextWithDot := eurText[0] + "." + eurText[1]
								currentEur, err = strconv.ParseFloat(euroTextWithDot, 64)
							}
						})
					}
				})
			}
		})

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(5 * time.Second)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
