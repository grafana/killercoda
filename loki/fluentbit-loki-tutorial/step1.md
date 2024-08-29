# Step 1: Environment setup

In this step, we will set up our environment by cloning the repository that contains our demo application and spinning up our observability stack using Docker Compose.

1. To get started, clone the repository that contains our demo application:

   ```bash
   git clone -b fluentbit-official  https://github.com/grafana/loki-fundamentals.git
   ```{{exec}}

1. Next we will spin up our observability stack using Docker Compose:

   ```bash
   docker-compose -f loki-fundamentals/docker-compose.yml up -d
   ```{{exec}}

   This will spin up the following services:

   ```console
   ✔ Container loki-fundamentals-grafana-1       Started                                                        
   ✔ Container loki-fundamentals-loki-1          Started                        
   ✔ Container loki-fundamentals_fluent-bit_1    Started
   ```{{copy}}

Once we have finished configuring the Fluent Bit agent and sending logs to Loki, we will be able to view the logs in Grafana. To check if Grafana is up and running, navigate to the following URL: [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
