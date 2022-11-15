#!/bin/bash

pushd /home/exporter/eur-currency-prometheus-exporter

kill -9 $(ps aux|grep eur-currency-prometheus-exporter | awk 'NR==1{print $2}') || true
git reset origin --hard
git pull
rm -rf ./eur-currency-prometheus-exporter
go build .
nohup ./eur-currency-prometheus-exporter &

popd +1