# Quickstart to run Loki locally

If you want to experiment with Loki, you can run Loki locally using the Docker Compose file that ships with Loki. It runs Loki in [monolithic mode](https://grafana.com/docs/loki/latest/get-started/deployment-modes/#monolithic-mode) and includes a sample application to generate logs. For production deployments, see [deployment modes](https://grafana.com/docs/loki/latest/get-started/deployment-modes/) for currently recommended options.

The Docker Compose configuration runs the following components, each in its own container:

- **flog**: which generates log lines.
  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.
- **Grafana Alloy**: which scrapes the log lines from flog, and pushes them to Loki through the gateway.
- **Gateway** (nginx) which receives requests and redirects them to the Loki container.
- **Loki**: which runs all Loki components (Distributor, Ingester, Query Frontend, Query Scheduler, Querier, Compactor, and Ruler) together in a single monolithic process.
- **Silo**: which Loki uses to store its index and chunks.
- **Grafana**: which provides visualization of the log lines captured within Loki.

![Getting started sample application](https://grafana.com/media/docs/loki/get-started-flog-v3.png)
