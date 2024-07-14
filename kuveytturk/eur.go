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

func setCurrentEur() {

	for {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://www.kuveytturk.com.tr/ck0d84?C24AD4C0FDA76C73081889B634A8C039", nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Cookie", "NSC_WTSW-XFCTJUF-TTM=ffffffffaf1c1ca845525d5f4f58455e445a4a42378b; TS01e71088=0176dcd71c45d828f601804554a1b3d6fe9b3c3039eb4cfc375a376e36d668b7912f2937bbc75a2f018d9a33dd6ffb1cd437388b626329fb18ea8d84a40431307dd6dc9c7e; _ga=GA1.3.1626640906.1720965814; _gid=GA1.3.482756998.1720965814")
		req.Header.Set("Referer", "https://www.kuveytturk.com.tr/finans-portali")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
		req.Header.Set("X-Bone-Language", "TR")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"macOS"`)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var currentEur float64
		value := gjson.Get(string(bodyText), "2.BuyRate")
		currentEur, err = strconv.ParseFloat(value.String(), 64)

		currentCurrencyMetric.WithLabelValues().Set(currentEur)
		time.Sleep(5 * time.Second)
	}
}

func CurrentEur() *prometheus.GaugeVec {
	go setCurrentEur()
	return currentCurrencyMetric
}
