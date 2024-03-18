# Step 4: Creating a Basic Dashboard
In this step, we will create a basic dashboard in Grafana to visualize the metrics collected by Prometheus. We will start by adding a new dashboard and then create a simple graph to display the CPU usage of our virtual environment.

## Adding a New Dashboard
1. Open your browser and go to the Grafana UI by clicking on the following link: **[http://localhost:3000]({{TRAFFIC_HOST1_3000}})**
2. Sign in with the following credentials:
   - **Username:** admin
   - **Password:** admin
**Note:** If you are already logged in, you can skip this step.
3. Once logged in, open the left hand side menu, then click on **Dashboards**. Next select the button ** + Create Dashboard**. 

## Creating a Graph Panel
4. Click on the **Add visualization** button to add a new panel to the dashboard.
5. Select **Prometheus** from the list of data sources.
6. Select the **Time Series** visualization from the right handside panel.
7. In the **Query** section, set the following fields:
   - **Data Source:** `Prometheus`
8. In the **Metrics** section, set the following fields:
   - **Query:** `100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)`
   - **Legend:** `CPU Usage`
   - **Unit:** `percent(0-100)`
   - **Min:** `0`
   - **Max:** `100`
9.  Click on the **Apply** button to apply the changes to the graph panel.
10. Click on the **Save dashboard** button to save the new dashboard.
11. Enter a name for the dashboard and click on the **Save** button.
