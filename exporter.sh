#!/bin/bash
pushd /home/ubuntu/eur-currency-prometheus-exporter

sudo kill -9 $(ps aux|grep eur-currency-prometheus-exporter | awk 'NR==1{print $2}') || true
nohup ./eur-currency-prometheus-exporter &

popd +1