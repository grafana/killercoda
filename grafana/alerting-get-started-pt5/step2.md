# Use case: monitoring and alerting for system health with Prometheus and Grafana

In this use case, we focus on monitoring the system's CPU, memory, and disk usage as part of a monitoring setup. This example is based on the [Grafana Prometheus Alerting Demo](https://github.com/tonypowa/grafana-prometheus-alerting-demo), which collects and visualizes system metrics via Prometheus and Grafana.

Your team is responsible for ensuring the health of your servers, and you want to leverage advanced alerting features in Grafana to:

- Set who should receive an alert notification based on query value.
- Suppress alerts based on query value.

## Scenario

In the provided demo setup, you're monitoring:

- CPU Usage.
- Memory Consumption.

You have a mixture of critical alerts (e.g., CPU usage over `75%`{{copy}}) and warning alerts (e.g., memory usage over `60%`{{copy}}).

This Flask-based Python script simulates a service that:

- Generates random CPU and memory usage values (10% to 100%) every **10 seconds**
- Exposes them as Prometheus metrics
- Each metric includes a default instance label based on the scrape target:
  - `instance="flask-prod:5000"`{{copy}}
  - `instance="flask-staging:5000"`{{copy}}
- A custom deployment label added explicitly in the app logic (this serves as an additional example for dynamically routing production instances):
  - `deployment="prod-us-cs30"`{{copy}}
  - `deployment="staging-us-cs20"`{{copy}}

## Objective

Use templates to dynamically populate a custom label that matches a notification policy, and therefore routes alerts to the correct contact point.

We'll automatically determine the environment associated with each firing alert by inspecting system metrics (e.g., CPU, memory) and extracting keywords using regular expressions with the Go templating language.
