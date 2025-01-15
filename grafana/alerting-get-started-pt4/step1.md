To demonstrate the observation of data using the Grafana stack, download and run the following files.

1. Clone the [tutorial environment repository](https://www.github.com/grafana/tutorial-environment).

   ```
   git clone https://github.com/grafana/tutorial-environment.git
   ```{{exec}}

1. Change to the directory where you cloned the repository:

   ```
   cd tutorial-environment
   ```{{exec}}

1. Run the Grafana stack:

   ```bash
   docker-compose up -d
   ```{{exec}}

   The first time you run `docker compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

   NOTE:

   If you already have Grafana, Loki, or Prometheus running on your system, you might see errors, because the Docker image is trying to use ports that your local installations are already using. If this is the case, stop the services, then run the command again.
