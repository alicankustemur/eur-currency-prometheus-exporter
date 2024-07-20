package tcmb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Base struct {
	XMLName  xml.Name    `xml:"Tarih_Date"`
	Currency []*Currency `xml:"Currency"`
}

type Currency struct {
	XMLName      xml.Name `xml:"Currency"`
	CurrencyName string   `xml:"CurrencyName"`
	ForexBuying  float64  `xml:"ForexBuying"`
}

var currentCurrencyMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "tcbm_currency",
	Help: "Currency Prices of TCMB",
}, []string{})

var currentEur float64

func setCurrentEur() {

	t := time.Now()

	for {

		dayStr := fmt.Sprintf("%d", t.Day())
		monthStr := fmt.Sprintf("%d", t.Month())

		if t.Day() <= 9 {
			dayStr = fmt.Sprintf("0%d", t.Day())
		}

		if t.Month() <= 9 {
			monthStr = fmt.Sprintf("0%d", t.Month())
		}

		url := fmt.Sprintf("https://www.tcmb.gov.tr/kurlar/%d%s/%s%s%d.xml", t.Year(), monthStr, dayStr, monthStr, t.Year())

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(3 * time.Second)

		res, err := http.DefaultClient.Do(req)

		if res.StatusCode == 404 {
			t = t.AddDate(0, 0, -1)
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		var result Base
		xml.Unmarshal(data, &result)

		for _, currency := range result.Currency {

			if currency.CurrencyName == "EURO" {
				currentEur = currency.ForexBuying
			}

		}

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(1 * time.Hour)
		t = time.Now()
	}

}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
