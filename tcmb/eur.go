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

	for {

		t := time.Now()

		day := time.Now().Day()

		url := fmt.Sprintf("https://www.tcmb.gov.tr/kurlar/%d%d/%d%d%d.xml", t.Year(), t.Month(), day, t.Month(), t.Year())

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)

		data, err := ioutil.ReadAll(res.Body)

		var result Base
		xml.Unmarshal(data, &result)

		for _, currency := range result.Currency {

			if currency.CurrencyName == "EURO" {
				currentEur = currency.ForexBuying
			}

		}

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(2 * time.Hour)
	}

}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
