# Before you begin

To demonstrate the observation of data using the Grafana stack, download and run the following files.

1. Create a folder for this tutorial, then download and save the Docker Compose file, the Alloy configuration, and the Grafana data source provisioning file for this tutorial. Also create a `logs`{{copy}} folder, which Alloy reads the sample logs from.

   ```bash
   mkdir alerts-with-logs && cd alerts-with-logs
   wget https://raw.githubusercontent.com/grafana/tutorial-environment/master/loki-logs/docker-compose.yaml
   wget https://raw.githubusercontent.com/grafana/tutorial-environment/master/loki-logs/config.alloy
   mkdir -p provisioning/datasources logs
   wget -P provisioning/datasources https://raw.githubusercontent.com/grafana/tutorial-environment/master/loki-logs/provisioning/datasources/datasources.yaml
   ```{{exec}}
1. Run the Grafana stack.

   ```bash
   docker-compose up -d
   ```{{exec}}

The first time you run `docker-compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

> If you already have Grafana, Loki, or Prometheus running on your system, you might see errors, because the Docker image is trying to use ports that your local installations are already using. If this is the case, stop the services, then run the command again.
