# Start Grafana Mimir and dependencies

Start running your local setup with the following Docker command:

```bash
docker-compose up -d
```{{exec}}

This command starts:

- Grafana Mimir
  - Three instances of monolithic-mode Mimir to provide high availability

  - Multi-tenancy enabled (tenant ID is `demo`{{copy}})

- [Minio](https://min.io/)
  - S3-compatible persistent storage for blocks, rules, and alerts

- Prometheus
  - Scrapes Grafana Mimir metrics, then writes them back to Grafana Mimir to ensure availability of ingested metrics

- Grafana
  - Includes a preinstalled datasource to query Grafana Mimir

  - Includes preinstalled dashboards for monitoring Grafana Mimir

- Load balancer
  - A simple NGINX-based load balancer that exposes Grafana Mimir endpoints on the host

The diagram below illustrates the relationship between these components:
![Architecture diagram for this Grafana Mimir tutorial](https://grafana.com/tutorial-architecture.png)

The following ports will be exposed on the host:

- Grafana on [`http://localhost:9000`{{copy}}]({{TRAFFIC_HOST1_9000}})

- Grafana Mimir on [`http://localhost:9009`{{copy}}]({{TRAFFIC_HOST1_9009}})

To learn more about the Grafana Mimir configuration, you can review the configuration file `config/mimir.yaml`{{copy}}.
