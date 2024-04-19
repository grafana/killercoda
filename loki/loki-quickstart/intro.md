<p><img src="../assets/loki-logo.png" style=" height: 90px"></p>


# Quickstart to run Loki locally

 If you want to experiment with Loki, you can run Loki locally using the Docker Compose file that ships with Loki. It runs Loki in a [monolithic deployment](https://grafana.com/docs/loki/latest/get-started/deployment-modes/#monolithic-mode) mode and includes a sample application to generate logs.

The Docker Compose configuration instantiates the following components, each in its own container:

- **flog** a sample application which generates log lines.  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.
- **Promtail** which scrapes the log lines from flog, and pushes them to Loki through the gateway.
- **Gateway** (NGINX) which receives requests and redirects them to the appropriate container based on the request's URL.
- One Loki **read** component.
- One Loki **write** component.
- **Minio** an S3-compatible object store which Loki uses to store its index and chunks.
- **Grafana** which provides visualization of the log lines captured within Loki.

![Loki Quickstart](https://grafana.com/media/docs/loki/get-started-flog-v2.png)
