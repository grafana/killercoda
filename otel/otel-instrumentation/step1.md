# Explore the LGTM Stack

The LGTM (Loki, Grafana, Tempo, Mimir) stack provides a complete observability solution:
- **Loki**: Log aggregation system
- **Grafana**: Visualization and monitoring platform
- **Tempo**: Distributed tracing backend
- **Mimir**: Metrics storage

We also have an OpenTelemetry Collector acting as a single gateway for our LGTM stack.

![LGTM stack](https://raw.githubusercontent.com/grafana/docker-otel-lgtm/main/img/overview.png)

> In this lab, you will see Prometheus and Mimir mentionned. Those 2 terms are equivalent in our current context. Mimir is a scalable version of Prometheus.

## Starting the Stack

You can verify the stack is running with:
```bash
docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Status}}"
```{{exec}}

Two services should be running. The stack includes:
- LGTM. A single container running:
   - Loki
   - Grafana, accessible at [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
   - Tempo
   - Mimir
   - An OTEL gateway to receive OTLP signals and dispatch them to the backends
- Alloy: an OpenTelemetry Collector distro

## Verifying the Setup

1. Open Grafana at [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
2. You'll be automatically logged in as admin
3. Navigate to Drilldown > Metrics/Logs/Traces to verify:
   - Prometheus is configured for metrics
   - Loki is configured for logs
   - Tempo is configured for traces

For now, we only have some metrics about our OpenTelemetry gateway. The rest is empty. Don't worry, in the next step, we'll explore how to use these tools to monitor our demo application.
