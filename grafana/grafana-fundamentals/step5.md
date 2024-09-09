# Explore your logs

Grafana Explore not only lets you make ad-hoc queries for metrics, but lets you explore your logs as well.

1. Click the menu icon and, in the sidebar, click **Explore**.

1. In the data source list at the top, select the **Loki** data source.

1. Confirm that you’re in code mode by checking the **Builder/Code** toggle at the top right corner of the query panel.

1. Enter the following in the query editor, and then press Shift + Enter:

   ```
   {filename="/var/log/tns-app.log"}
   ```{{copy}}

1. Grafana displays all logs within the log file of the sample application. The height of each bar in the graph encodes the number of logs that were generated at that time.

1. Click and drag across the bars in the graph to filter logs based on time.

Not only does Loki let you filter logs based on labels, but on specific occurrences.

Let’s generate an error, and analyze it with Explore.

1. In the [sample application]({{TRAFFIC_HOST1_8081}}), post a new link without a URL to generate an error in your browser that says `empty url`{{copy}}.

1. Go back to Grafana and enter the following query to filter log lines based on a substring:

   ```
   {filename="/var/log/tns-app.log"} |= "error"
   ```{{copy}}

1. Click the log line that says `level=error msg="empty url"`{{copy}} to see more information about the error.

   > **Note:** If you’re in Live mode, clicking logs will not show more information about the error. Instead, stop and exit the live stream, then click the log line there.

Logs are helpful for understanding what went wrong. Later in this tutorial, you’ll see how you can correlate logs with metrics from Prometheus to better understand the context of the error.
