# Install Loki and collecting sample logs

**To install Loki locally, follow these steps:**

1. Create a directory called `evaluate-loki`{{copy}} for the demo environment.
   Make `evaluate-loki`{{copy}} your current working directory:

   ```bash
   mkdir evaluate-loki
   cd evaluate-loki
   ```{{exec}}
1. Download `loki-config.yaml`{{copy}}, `alloy-local-config.yaml`{{copy}}, and `docker-compose.yaml`{{copy}}:

   ```bash
   wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/loki-config.yaml -O loki-config.yaml
   wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/alloy-local-config.yaml -O alloy-local-config.yaml
   wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/docker-compose.yaml -O docker-compose.yaml
   ```{{exec}}
1. Deploy the sample Docker image.

   With `evaluate-loki`{{copy}} as the current working directory, start the demo environment using `docker compose`{{copy}}:

   ```bash
   docker compose up -d
   ```{{exec}}

   At the end of the command, you should see something similar to the following:

   ```console
   ✔ Network evaluate-loki_loki          Created      0.1s
   ✔ Container evaluate-loki_silo_1      Started      0.6s
   ✔ Container evaluate-loki_flog_1      Started      0.6s
   ✔ Container evaluate-loki_loki_1      Started      0.8s
   ✔ Container evaluate-loki_gateway_1   Started      1.1s
   ✔ Container evaluate-loki_grafana_1   Started      1.4s
   ✔ Container evaluate-loki_alloy_1     Started      1.4s
   ```{{copy}}
1. (Optional) Verify that the Loki cluster is up and running.

   - The Loki component returns `ready`{{copy}} when you browse to [http://localhost:3101/ready]({{TRAFFIC_HOST1_3101}}/ready).
     It can take a few seconds after the container starts for all internal components to report ready.
1. (Optional) Verify that Grafana Alloy is running.

   - You can access the Grafana Alloy UI at [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).
1. (Optional) You can check all the containers are running by running the following command:

   ```bash
   docker ps -a
   ```{{exec}}
