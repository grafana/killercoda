# Loki Tutorial

This quick start guide will walk you through deploying Loki in single binary mode (also known as [monolithic mode](https://grafana.com/docs/loki/latest/get-started/deployment-modes/#monolithic-mode)) using Docker Compose. Grafana Loki is only one component of the Grafana observability stack for logs. In this tutorial we will refer to this stack as the **Loki stack**.

![Loki stack](https://grafana.com/media/docs/loki/getting-started-loki-stack-3.png)

The Loki stack consists of the following components:

- **Alloy**: [Grafana Alloy](https://grafana.com/docs/alloy/latest/) is an open source telemetry collector for metrics, logs, traces, and continuous profiles. In this quickstart guide Grafana Alloy has been configured to tail logs from all Docker containers and forward them to Loki.

- **Loki**: A log aggregation system to store the collected logs. For more information on what Loki is, see the [Loki overview](https://grafana.com/docs/loki/latest/get-started/overview/).

- **Grafana**: [Grafana](https://grafana.com/docs/grafana/latest/) is an open-source platform for monitoring and observability. Grafana will be used to query and visualize the logs stored in Loki.
