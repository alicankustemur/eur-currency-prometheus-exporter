module github.com/alicankustemur/eur-currency-prometheus-exporter

go 1.18

require (
	github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom v0.0.0-00010101000000-000000000000
	github.com/alicankustemur/eur-currency-prometheus-exporter/kuveytturk v0.0.0-00010101000000-000000000000
	github.com/alicankustemur/eur-currency-prometheus-exporter/tcmb v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.15.1
)

require (
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom => ./dovizcom

replace github.com/alicankustemur/eur-currency-prometheus-exporter/tcmb => ./tcmb

replace github.com/alicankustemur/eur-currency-prometheus-exporter/kuveytturk => ./kuveytturk

replace github.com/alicankustemur/eur-currency-prometheus-exporter/canlidoviz => ./canlidoviz