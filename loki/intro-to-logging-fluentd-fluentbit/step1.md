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
   ✔ Container loki-fundamentals_grafana_1  Started                                                        
   ✔ Container loki-fundamentals_loki_1     Started                        
   ✔ Container loki-fundamentals_fluentd_1  Started
   ✔ Container loki-fundamentals_fluentbit_1 Started
   ```{{copy}}
