# Deploy the Loki stack

**To deploy the Loki stack locally, follow these steps:**

1. Clone the Loki fundamentals repository and check out the getting-started branch:

   ```bash
   git clone https://github.com/grafana/loki-fundamentals.git -b getting-started
   ```{{exec}}
1. Change to the `loki-fundamentals`{{copy}} directory:

   ```bash
   cd loki-fundamentals
   ```{{exec}}
1. With `loki-fundamentals`{{copy}} as the current working directory deploy Loki, Alloy, and Grafana using Docker Compose:

   ```bash
   docker compose up -d
   ```{{exec}}

   After running the command, you should see a similar output:

   ```console
    ✔ Container loki-fundamentals-grafana-1  Started  0.3s 
    ✔ Container loki-fundamentals-loki-1     Started  0.3s 
    ✔ Container loki-fundamentals-alloy-1    Started  0.4s
   ```{{copy}}
1. With the Loki stack running, you can now verify each component is up and running:

   - **Alloy**: Open a browser and navigate to [http://localhost:12345/graph]({{TRAFFIC_HOST1_12345}}/graph). You should see the Alloy UI.
   - **Grafana**: Open a browser and navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}). You should see the Grafana home page.
   - **Loki**: Open a browser and navigate to [http://localhost:3100/metrics]({{TRAFFIC_HOST1_3100}}/metrics). You should see the Loki metrics page.
