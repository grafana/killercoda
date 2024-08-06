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

   This will spin up the following services:

   ```console
   ✔ Container loki-fundamentals-grafana-1          Started                                                        
   ✔ Container loki-fundamentals-loki-1             Started                        
   ✔ Container loki-fundamentals_otel-collector_1   Started
   ```{{copy}}

   **Note:** The OpenTelemetry Collector container will show as `Stopped`{{copy}}. This is expected as we have provided an empty configuration file. We will update this file in the next step.

Once we have finished configuring the OpenTelemetry Collector and sending logs to Loki, we will be able to view the logs in Grafana. To check if Grafana is up and running, navigate to the following URL: [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
