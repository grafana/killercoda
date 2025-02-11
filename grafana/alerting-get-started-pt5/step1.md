# Set up the Grafana stack

To demonstrate the observation of data using the Grafana stack, download and run the following files.

1. Clone the [tutorial environment repository](https://github.com/tonypowa/grafana-prometheus-alerting-demo.git).

   ```bash
   git clone https://github.com/tonypowa/grafana-prometheus-alerting-demo.git
   ```{{exec}}

1. Change to the directory where you cloned the repository:

   ```bash
   cd grafana-prometheus-alerting-demo
   ```{{exec}}

1. Build the Grafana stack:

   ```bash
   docker-compose build
   ```{{exec}}

1. Bring up the containers:

   ```bash
   docker-compose up -d
   ```{{exec}}

   The first time you run `docker compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

NOTE:

If you already have Grafana, Loki, or Prometheus running on your system, you might see errors, because the Docker image is trying to use ports that your local installations are already using. If this is the case, stop the services, then run the command again.
