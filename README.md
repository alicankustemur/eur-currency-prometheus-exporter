# Setup exporter

clone the repo and build the project
```
git clone https://github.com/alicankustemur/eur-currency-prometheus-exporter
cd eur-currency-prometheus-exporter
go build .
```

Create an `exporter.sh` file
```
sudo vi /usr/bin/exporter.sh
# /usr/bin/exporter.sh
#!/bin/bash
pushd /home/ubuntu/eur-currency-prometheus-exporter

sudo kill -9 $(ps aux|grep eur-currency-prometheus-exporter | awk 'NR==1{print $2}') || true
nohup ./eur-currency-prometheus-exporter &

popd +1
```
Give an executable permission to `exporter.sh`.
```
sudo chmod +x /usr/bin/exporter.sh
```

then create an exporter systemd service.

```
sudo vi /etc/systemd/system/exporter.service

[Unit]
Description=Prometheus Exporter
After=network.target
StartLimitIntervalSec=0

[Service]
Type=forking
Restart=always
RestartSec=1
User=root
Group=root
ExecStart=/usr/bin/exporter.sh

[Install]
WantedBy=multi-user.target
```

```
sudo systemctl start exporter.service
```

the service must be on `active` state
```
sudo systemctl status exporter.service

● exporter.service - Prometheus Exporter
     Loaded: loaded (/etc/systemd/system/exporter.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2023-09-27 16:35:11 UTC; 22min ago
    Process: 759 ExecStart=/usr/bin/exporter.sh (code=exited, status=0/SUCCESS)
   Main PID: 835 (eur-currency-pr)
      Tasks: 7 (limit: 1060)
     Memory: 45.1M
        CPU: 2.680s
     CGroup: /system.slice/exporter.service
             └─835 ./eur-currency-prometheus-exporter

Sep 27 16:35:10 exporter systemd[1]: Starting Prometheus Exporter...
Sep 27 16:35:10 exporter exporter.sh[759]: /home/ubuntu/eur-currency-prometheus-exporter /
Sep 27 16:35:11 exporter sudo[773]:     root : PWD=/home/ubuntu/eur-currency-prometheus-exporter ; USER=root ; COMMAND=/usr/bin/kill -9 765
Sep 27 16:35:11 exporter sudo[773]: pam_unix(sudo:session): session opened for user root(uid=0) by (uid=0)
Sep 27 16:35:11 exporter exporter.sh[834]: kill: (765): No such process
Sep 27 16:35:11 exporter sudo[773]: pam_unix(sudo:session): session closed for user root
Sep 27 16:35:11 exporter exporter.sh[759]: /home/ubuntu/eur-currency-prometheus-exporter
Sep 27 16:35:11 exporter systemd[1]: Started Prometheus Exporter.
```

let's enable the service at system boot

```
sudo systemctl enable exporter.service
```
