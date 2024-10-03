# Step 1: Environment setup

In this step, we will set up our environment by cloning the repository that contains our demo application and spinning up our observability stack using Docker Compose.

1. To get started, clone the repository that contains our demo application:

   ```bash
   git clone -b microservice-otel-collector  https://github.com/grafana/loki-fundamentals.git
   ```{{exec}}

1. Next we will spin up our observability stack using Docker Compose:

   ```bash
   docker-compose -f loki-fundamentals/docker-compose.yml up -d 
   ```{{exec}}

   To check the status of services we can run the following command:

   ```bash
   docker ps -a
   ```{{exec}}

After we’ve finished configuring the OpenTelemetry Collector and sending logs to Loki, we will be able to view the logs in Grafana. To check if Grafana is up and running, navigate to the following URL: [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
