# Use case: Monitoring and alerting for system health with Prometheus and Grafana

In this use case, we focus on monitoring the system's CPU, memory, and disk usage as part of a monitoring setup. This example is based on the [Grafana Prometheus Alerting Demo](https://github.com/tonypowa/grafana-prometheus-alerting-demo), which collects and visualizes system metrics via Prometheus and Grafana.

Your team is responsible for ensuring the health of your servers, and you want to leverage advanced alerting features in Grafana to:

- Set who should receive an alert notification based on query value.
- Suppress alerts based on query value.
- Integrate alert rules into visualizations for better context.

## Scenario

In the provided demo setup, you're monitoring:

- CPU Usage.
- Memory Consumption.

You have a mixture of critical alerts (e.g., CPU usage over `75%`{{copy}}) and warning alerts (e.g., memory usage over `60%`{{copy}}).

At times, you also have scheduled maintenance windows, where you might temporarily suppress certain alerts during planned downtime.
