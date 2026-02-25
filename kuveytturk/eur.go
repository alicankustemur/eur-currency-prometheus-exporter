package kuveytturk

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

var currentCurrencyMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "kuveytturk_currency",
	Help: "Current Currency Price of Kuveyt Turk",
}, []string{})

var httpClient = &http.Client{Timeout: 30 * time.Second}

func setCurrentEur() {
	for {
		req, err := http.NewRequest("GET", "https://www.kuveytturk.com.tr/ck0d84?C24AD4C0FDA76C73081889B634A8C039", nil)
		if err != nil {
			log.Println("kuveytturk: error creating request:", err)
			time.Sleep(1 * time.Minute)
			continue
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Referer", "https://www.kuveytturk.com.tr/finans-portali")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
		req.Header.Set("X-Bone-Language", "TR")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Println("kuveytturk: error fetching:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		bodyText, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Println("kuveytturk: error reading body:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		value := gjson.Get(string(bodyText), "2.BuyRate")
		currentEur, err := strconv.ParseFloat(value.String(), 64)
		if err != nil {
			log.Println("kuveytturk: error parsing rate:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(1 * time.Minute)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
