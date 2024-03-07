# Step 1: Collecting Metrics
Before we can start visualizing our data, we need to collect some metrics. The easiest todo this is by collecting our virtiual enviroments system metrics. To do this we will use Node Exporter and Prometheus. Lets break down each of these components and we can install them within our virtual enviroment.

## Node Exporter
Node Exporter is a Prometheus exporter for hardware and OS metrics exposed by *nix kernels, written in Go with pluggable metric collectors. It allows for the collection of hardware and OS metrics and is a great way to collect system metrics for your virtual enviroment. Lets install Node Exporter on our virtual enviroment:

1. Download the latest version of Node Exporter from the [Prometheus download page](https://prometheus.io/download/).
```
wget https://github.com/prometheus/node_exporter/releases/download/v1.7.0/node_exporter-1.7.0.linux-amd64.tar.gz
```{{exec}}

2. Extract the tarball and navigate to the extracted directory. Move to bin directory.
```
tar -xvf node_exporter-1.7.0.linux-amd64.tar.gz && cd node_exporter-1.7.0.linux-amd64 && sudo cp node_exporter /usr/local/bin
```{{exec}}

3. Lets turn Node Exporter into a service (we will do the hard work for you).
```
sudo useradd -rs /bin/false node_exporter && echo -e "[Unit]\nDescription=Node Exporter\nAfter=network.target\n\n[Service]\nUser=node_exporter\nGroup=node_exporter\nType=simple\nExecStart=/usr/local/bin/node_exporter\n\n[Install]\nWantedBy=multi-user.target" | sudo tee /etc/systemd/system/node_exporter.service > /dev/null && sudo systemctl daemon-reload && sudo systemctl start node_exporter && sudo systemctl status node_exporter

```
