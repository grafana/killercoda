# Log into Grafana and explore metrics in Prometheus

Open [http://localhost:3000/explore/metrics/]({{TRAFFIC_HOST1_3000}}/explore/metrics/) to access the **Explore Metrics** feature in Grafana.

From here you can visually explore the metrics that are being sent to Prometheus by Alloy.

![Explore Metrics App](https://grafana.com/media/docs/alloy/explore-metrics.png)

You can also build promQL queries manually to explore the data further.

Open [http://localhost:3000/explore]({{TRAFFIC_HOST1_3000}}/explore) to access the **Explore** feature in Grafana.

Select Prometheus as the data source and click the **Metrics Browser** button to select the metric, labels, and values for your labels.

Here you can see that metrics are flowing through to Prometheus as expected, and the end-to-end configuration was successful.

{{< figure src="/media/docs/alloy/tutorial/Metrics_visualization.png" alt=“Your data flowing through Prometheus.” >}}
