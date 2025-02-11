# Log in to Grafana and explore Loki logs

Open [http://localhost:3000/explore]({{TRAFFIC_HOST1_3000}}/explore) to access **Explore** feature in Grafana.

Select Loki as the data source and click the **Label Browser** button to select a file that Alloy has sent to Loki.

Here you can see that logs are flowing through to Loki as expected, and the end-to-end configuration was successful.

![Logs reported by Alloy in Grafana](https://grafana.com/media/docs/alloy/tutorial/loki-logs.png)
