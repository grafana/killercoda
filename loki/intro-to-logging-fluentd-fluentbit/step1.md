# Step 1: Environment setup

In this step, we will set up our environment by cloning the repository that contains our demo application and spinning up our observability stack using Docker Compose.

1. To get started, clone the repository that contains our demo application:

   ```bash
   git clone -b microservice-fluentd-fluentbit  https://github.com/grafana/loki-fundamentals.git
   ```{{exec}}

1. Next we will spin up our observability stack using Docker Compose:

   ```bash
   docker-compose -f loki-fundamentals/docker-compose.yml up -d 
   ```{{exec}}

   This will spin up the following services:

   ```bash
   ✔ Container loki-fundamentals-grafana-1  Started                                                        
   ✔ Container loki-fundamentals-loki-1     Started                        
   ✔ Container loki-fundamentals-fluentd-1  Started
   ✔ Container loki-fundamentals-fluentbit-1 Started
   ```{{copy}}
