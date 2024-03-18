# Step 4: Creating a Basic Dashboard
In this step, we will create a basic dashboard in Grafana to visualize the metrics collected by Prometheus. We will start by adding a new dashboard and then create a simple graph to display the CPU usage of our virtual environment.

## Adding a New Dashboard
1. Open your browser and go to the Grafana UI by clicking on the following link: **[http://localhost:3000]({{TRAFFIC_HOST1_3000}})**
2. Sign in with the following credentials:
   - **Username:** admin
   - **Password:** admin

**Note:** If you are already logged in, you can skip this step.

3. Once logged in, open the left hand side menu, then click on **Dashboards**. Next select the button **+ Create Dashboard**. 

## Creating a Graph Panel
1. Click on the **Add visualization** button to add a new panel to the dashboard.
2. Select **Prometheus** from the list of data sources.
3. Select the **Time Series** visualization from the right handside panel.
4. In the **Query** section, set the following fields:
   - **Data Source:** `Prometheus`
5. In the **Query** tab, toggle from **Builder** to **Code** and enter the following promQL query:
   - **Query:** `100 - (avg by (instance) (irate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)`
  
*This query calculates the average CPU utilization (not in idle mode) for each instance over the past 5 minutes. It subtracts the average percentage of time each CPU instance has been idle (as measured over 5-minute intervals) from 100%, effectively giving the average active CPU usage.*

6. Click on the **Run** button to execute the query and visualize the results.
8. Lets add some additional configuration to our time series graph:
   - **Title:** `CPU Usage`
   - **Unit:** `percent(0-100)`
   - **Min:** `0`
   - **Max:** `100`

*Top Tip: these are all configured via the configuration panel on the right handside. You can make use of the search bar within this panel to access the above fields quickly*

1.  Click on the **Apply** button to apply the changes to the graph panel.
2.  Click on the **Save dashboard** button to save the new dashboard.
3.  Enter a name for the dashboard and click on the **Save** button.
