# Step 1: Environment setup

In this step, we will set up our environment by cloning the repository that contains our demo application and spinning up our observability stack using Docker Compose.

1. To get started, clone the repository that contains our demo application:

   ```bash
   git clone -b microservice-otel  https://github.com/grafana/loki-fundamentals.git
   ```{{exec}}

1. Next we will spin up our observability stack using Docker Compose:

   ```bash
   docker-compose -f loki-fundamentals/docker-compose.yml up -d 
   ```{{exec}}

   This will spin up the following services:

   ```console
   ✔ Container loki-fundamentals-grafana-1  Started                                                        
   ✔ Container loki-fundamentals-loki-1     Started                        
   ✔ Container loki-fundamentals-alloy-1    Started
   ```{{copy}}

We will be access two UI interfaces:

- Alloy at [http://localhost:12345]({{TRAFFIC_HOST1_12345}})

- Grafana at [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
