# Before you begin

To demonstrate the observation of data using the Grafana stack, download the files to your local machine.

1. Download and save a Docker compose file to run Grafana, Loki and Promtail.

   ```bash
   wget https://raw.githubusercontent.com/grafana/loki/v2.8.0/production/docker-compose.yaml -O docker-compose.yaml
   ```{{exec}}

1. Run the Grafana stack.

   ```bash
   docker-compose up -d
   ```{{exec}}

The first time you run `docker-compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

> If you already have Grafana, Loki, or Prometheus running on your system, you might see errors, because the Docker image is trying to use ports that your local installations are already using. If this is the case, stop the services, then run the command again.
