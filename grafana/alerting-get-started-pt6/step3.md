# Step 1: Create a visualization to monitor metrics

To keep track of these metrics you can set up a visualization for CPU usage and memory consumption. This will make it easier to see how the system is performing.

The time-series visualization supports alert rules to provide more context in the form of annotations and alert rule state. Follow these steps to create a visualization to monitor the application’s metrics.

1. Log in to Grafana:

   - Navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana should be running.
   - Username and password: `admin`{{copy}}
1. Create a time series panel:

   - Navigate to **Dashboards**.
   - Click **+ Create dashboard**.
   - Click **+ Add visualization**.
   - Select **Prometheus** as the data source (provided with the demo).
   - Enter a title for your panel, e.g., **CPU and Memory Usage**.
1. Add queries for metrics:

   - In the query area, copy and paste the following PromQL query:

     ** switch to **Code** mode if not already selected **

     ```promql
     flask_app_cpu_usage{instance="flask-prod:5000"}
     ```{{copy}}
   - Click **Run queries**.

   This query should display the simulated CPU usage data for the **prod** environment.
1. Add memory usage query:

   - Click **+ Add query**.
   - In the query area, paste the following PromQL query:

     ```promql
     flask_app_memory_usage{instance="flask-prod:5000"}
     ```{{copy}}

   ![Time-series panel displaying CPU and memory usage metrics in production.](https://grafana.com/media/docs/alerting/cpu-mem-dash.png)
1. Click **Save dashboard**. Name it: `cpu-and-memory-metrics`{{copy}}.

We have our time-series panel ready. Feel free to combine metrics with labels such as `flask_app_cpu_usage{instance=“flask-staging:5000”}`{{copy}}, or other labels like `deployment`{{copy}}.
