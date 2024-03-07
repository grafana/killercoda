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
    ```{{exec}}

## Prometheus
Prometheus is a monitoring and alerting toolkit that is designed for reliability, scalability, and maintainability. It is a powerful tool for collecting and querying metrics and is a great way to store the metrics collected by Node Exporter. Lets install Prometheus on our virtual enviroment:

1. Download the latest version of Prometheus from the [Prometheus download page](https://prometheus.io/download/).
    ```
    wget https://github.com/prometheus/prometheus/releases/download/v2.50.1/prometheus-2.50.1.linux-amd64.tar.gz
    ```{{exec}}

2. Extract the tarball and navigate to the extracted directory. Move to bin directory.
    ```
    tar -xvf prometheus-2.50.1.linux-amd64.tar.gz && cd prometheus-2.50.1.linux-amd64 && sudo cp prometheus /usr/local/bin
    ```{{exec}}

3. Lets turn Prometheus into a service (we will do the hard work for you).
    ```
    sudo useradd -rs /bin/false prometheus && sudo mkdir /etc/prometheus /var/lib/prometheus && sudo cp prometheus.yml /etc/prometheus/prometheus.yml && sudo cp -r consoles/ console_libraries/ /etc/prometheus/ && sudo chown -R prometheus:prometheus /etc/prometheus /var/lib/prometheus && echo -e "[Unit]\nDescription=Prometheus\nAfter=network.target\n\n[Service]\nUser=prometheus\nGroup=prometheus\nType=simple\nExecStart=/usr/local/bin/prometheus --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/var/lib/prometheus --web.console.templates=/etc/prometheus/consoles --web.console.libraries=/etc/prometheus/console_libraries\n\n[Install]\nWantedBy=multi-user.target" | sudo tee /etc/systemd/system/prometheus.service > /dev/null && sudo systemctl daemon-reload && sudo systemctl start prometheus && sudo systemctl status prometheus
    ```{{exec}}
