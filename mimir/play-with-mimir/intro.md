# Play with Mimir

Grafana Mimir is a distributed, horizontally scalable, and highly available long term storage for [Prometheus](https://prometheus.io).

In this tutorial, you'll:

- Run Grafana Mimir locally with Docker Compose
- Run Prometheus to scrape some metrics and remote write to Grafana Mimir
- Run Grafana to explore Grafana Mimir dashboards
- Configure a testing recording rule and alert in Grafana Mimir

> **Note:**
> This tutorial runs Grafana Mimir using classic architecture. For more information about the supported architectures in Grafana Mimir, refer to [Grafana Mimir architecture](https://grafana.com/docs/mimir/latest/get-started/about-grafana-mimir-architecture/).
