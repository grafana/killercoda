# Explore your metrics

Grafana Explore is a workflow for troubleshooting and data exploration. In this step, you’ll be using Explore to create ad-hoc queries to understand the metrics exposed by the sample application.

> Ad-hoc queries are queries that are made interactively, with the purpose of exploring data. An ad-hoc query is commonly followed by another, more specific query.
1. Click the menu icon and, in the sidebar, click **Explore**. A dropdown menu for the list of available data sources is on the upper-left side. The Prometheus data source will already be selected. If not, choose Prometheus.

1. Confirm that you’re in code mode by checking the **Builder/Code** toggle at the top right corner of the query panel.

1. In the query editor, where it says _Enter a PromQL query…_, enter `tns_request_duration_seconds_count`{{copy}} and then press Shift + Enter.
   A graph appears.

1. In the top right corner, click the dropdown arrow on the **Run Query** button, and then select **5s**. Grafana runs your query and updates the graph every 5 seconds.

   You just made your first _PromQL_ query! [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) is a powerful query language that lets you select and aggregate time series data stored in Prometheus.

   `tns_request_duration_seconds_count`{{copy}} is a _counter_, a type of metric whose value only ever increases. Rather than visualizing the actual value, you can use counters to calculate the _rate of change_, i.e. how fast the value increases.

1. Add the [`rate`{{copy}}](https://prometheus.io/docs/prometheus/latest/querying/functions/#rate) function to your query to visualize the rate of requests per second. Enter the following in the query editor and then press Shift + Enter.

   ```
   rate(tns_request_duration_seconds_count[5m])
   ```{{copy}}

   Immediately below the graph there’s an area where each time series is listed with a colored icon next to it. This area is called the _legend_.

   PromQL lets you group the time series by their labels, using the [`sum`{{copy}}](https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators) aggregation operator.

1. Add the `sum`{{copy}} aggregation operator to your query to group time series by route:

   ```
   sum(rate(tns_request_duration_seconds_count[5m])) by(route)
   ```{{copy}}

1. Go back to the [sample application]({{TRAFFIC_HOST1_8081}}) and generate some traffic by adding new links, voting, or just refresh the browser.

1. Back in Grafana, in the upper-right corner, click the _time picker_, and select **Last 5 minutes**. By zooming in on the last few minutes, it’s easier to see when you receive new data.

Depending on your use case, you might want to group on other labels. Try grouping by other labels, such as `status_code`{{copy}}, by changing the `by(route)`{{copy}} part of the query to `by(status_code)`{{copy}}.
