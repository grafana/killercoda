# Use case: monitoring and alerting for system health with Prometheus and Grafana

In this use case, we focus on monitoring the system's CPU, memory, and disk usage as part of a monitoring setup. The [demo app](https://github.com/tonypowa/grafana-prometheus-alerting-demo), launches a stack that includes a Python script to simulate metrics, which Grafana collects and visualizes as a time-series visualization.

The script simulates random CPU and memory usage values (10% to 100%) every **10 seconds** and exposes them as Prometheus metrics.

## Objective

You'll build a time series visualization to monitor CPU and memory usage, define alert rules with threshold-based conditions, and link those alerts to your dashboards to display real-time annotations when thresholds are breached.
