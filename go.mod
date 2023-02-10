module github.com/alicankustemur/eur-currency-prometheus-exporter

go 1.18

require (
	github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom v0.0.0-00010101000000-000000000000
	github.com/alicankustemur/eur-currency-prometheus-exporter/enpara v0.0.0-00010101000000-000000000000
	github.com/alicankustemur/eur-currency-prometheus-exporter/kuveytturk v0.0.0-00010101000000-000000000000
	github.com/alicankustemur/eur-currency-prometheus-exporter/tcmb v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.14.0
)

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/alicankustemur/eur-currency-prometheus-exporter/enpara => ./enpara

replace github.com/alicankustemur/eur-currency-prometheus-exporter/tcmb => ./tcmb

replace github.com/alicankustemur/eur-currency-prometheus-exporter/dovizcom => ./dovizcom

replace github.com/alicankustemur/eur-currency-prometheus-exporter/kuveytturk => ./kuveytturk
