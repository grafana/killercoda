# Create a visualization to monitor metrics

To keep track of these metrics and understand system behavior across different environments, you can set up a visualization for CPU usage and memory consumption. This will make it easier to see how the system is performing and how alerts are distributed based on the environment label, including during scheduled maintenance windows.

The time-series visualization supports alert rules to provide more context in the form of annotations and alert rule state. Follow these steps to create a visualization to monitor the application’s metrics.

1. Log in to Grafana:

   Navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana should be running.

1. Create a time series panel:

   - Navigate to **Dashboards**.

   - Click **New**.

   - Select **New Dashboard**.

   - Click **+ Add visualization**.

   - Select **Prometheus** as the data source (provided with the demo).

   - Enter a title for your panel, e.g., **CPU and Memory Usage**.

1. Add queries for metrics:

   - In the query area, copy and paste the following PromQL query:

     ```promql
     flask_app_cpu_usage{environment="prod"}
     ```{{copy}}

   - Click **Run queries**.

   This query should display the simulated CPU usage data in the **prod** environment.

1. Add memory usage query:

   - Click **+ Add query**.

   - In the query area, paste the following PromQL query:

     ```promql
     flask_app_memory_usage{environment="prod"}
     ```{{copy}}

     ![Time-series panel displaying CPU and memory usage metrics in production.](https://grafana.com/media/docs/alerting/time-series_cpu_mem_usage_metrics.png)

     Both metrics return labels that we’ll use later to link alert instances with the appropriate routing. These labels help define how alerts are routed based on their environment or other criteria.

1. Click Save dashboard.
   We have our time-series panel ready. Feel free to combine metrics with labels such as `environment = “staging”`{{copy}}.
