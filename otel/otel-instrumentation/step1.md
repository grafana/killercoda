# Explore the LGTM Stack

> Give some time for your stack to start. You should see the images being downloaded in the console on the right. Wait for the task to complete.

The LGTM (Loki, Grafana, Tempo, Mimir) stack provides a complete observability solution:
- **Loki**: Log aggregation system
- **Grafana**: Visualization and monitoring platform
- **Tempo**: Distributed tracing backend
- **Mimir**: Metrics storage

## Starting the Stack

The environment initialization script has already:
1. Installed the Docker Compose plugin
2. Copied the necessary configuration files
3. Started the LGTM stack

You can verify the stack is running with:
```bash
docker-compose ps
```{{exec}}

All services should be in a healthy state (except permission-init, don't worry about it). The stack includes:
- Loki
- Grafana, accessible at [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
- Tempo
- Mimir

## Verifying the Setup

1. Open Grafana at [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
2. You'll be automatically logged in as admin
3. Navigate to Explore > Metrics/Logs/Traces to verify:
   - Mimir is configured for metrics
   - Loki is configured for logs
   - Tempo is configured for traces

> Mimir is 100% compatible with Prometheus. There is no "Mimir" datasource in Grafana, and you will see the Prometheus logo next to it.

For now, it's empty. Don't worry, in the next step, we'll explore how to use these tools to monitor our demo application.
