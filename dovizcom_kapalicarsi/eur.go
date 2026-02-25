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

var httpClient = &http.Client{Timeout: 30 * time.Second}

func setCurrentEur() {
	for {
		req, err := http.NewRequest("GET", "https://kur.doviz.com/kapalicarsi", nil)
		if err != nil {
			log.Println("dovizcom_kapalicarsi: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("dovizcom_kapalicarsi: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println("dovizcom_kapalicarsi: error parsing html:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		doc.Find("div.table,div.sortable").Each(func(i int, div *goquery.Selection) {
			if i == 0 {
				div.Find("tr").Each(func(j int, tr *goquery.Selection) {
					if j == 2 {
						tr.Find("td").Each(func(k int, td *goquery.Selection) {
							if k == 1 {
								eurText := strings.Split(td.Text(), ",")
								if len(eurText) >= 2 {
									euroTextWithDot := eurText[0] + "." + eurText[1]
									currentEur, err := strconv.ParseFloat(euroTextWithDot, 64)
									if err == nil {
										currentCurrencyMetric.WithLabelValues().Set(currentEur)
									}
								}
							}
						})
					}
				})
			}
		})

		time.Sleep(1 * time.Minute)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
