package tcmb

import (
	"encoding/xml"
	"fmt"
	"io"
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

var httpClient = &http.Client{Timeout: 30 * time.Second}

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
			log.Println("tcmb: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Println("tcmb: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		if res.StatusCode == 404 {
			res.Body.Close()
			t = t.AddDate(0, 0, -1)
			time.Sleep(3 * time.Second)
			continue
		}

		data, err := io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println("tcmb: error reading body:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		var result Base
		xml.Unmarshal(data, &result)

		for _, currency := range result.Currency {
			if currency.CurrencyName == "EURO" {
				currentCurrencyMetric.WithLabelValues().Set(currency.ForexBuying)
			}
		}

		time.Sleep(1 * time.Hour)
		t = time.Now()
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
