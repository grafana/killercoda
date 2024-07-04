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

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   At the end of the command, you should see something similar to the following:

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

   <!-- raw HTML omitted -->

1. (Optional) Verify that the Loki cluster is up and running.

   - The read component returns `ready`{{copy}} when you browse to [http://localhost:3101/ready]({{TRAFFIC_HOST1_3101}}/ready).
     The message `Query Frontend not ready: not ready: number of schedulers this worker is connected to is 0`{{copy}} shows until the read component is ready.

   - The write component returns `ready`{{copy}} when you browse to [http://localhost:3102/ready]({{TRAFFIC_HOST1_3102}}/ready).
     The message `Ingester not ready: waiting for 15s after being ready`{{copy}} shows until the write component is ready.

1. (Optional) Verify that Grafana Alloy is running.

   - You can access the Grafana Alloy UI at [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).
