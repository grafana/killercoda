# Clone the Tempo repository and start Docker

This quick start guide uses the `single-binary`{{copy}} docker-compose example. Check out our [collection of examples](https://github.com/grafana/tempo/blob/main/example/docker-compose) to explore various
operating modes and configurations.

Tempo requires a Kafka-compatible system for its write path, even when running as a single binary.
This example uses [Redpanda](https://redpanda.com/), a Kafka-compatible streaming platform, to handle trace ingestion.
Redpanda provides the durable queue that Tempo uses to decouple its write and read paths.

1. Clone the Tempo repository:

   ```bash
   git clone https://github.com/grafana/tempo.git
   ```{{exec}}
1. Go into the examples directory:

   ```bash
   cd tempo/example/docker-compose/single-binary
   ```{{exec}}
1. Start the services defined in the docker-compose file:

   ```bash
   docker compose up -d
   ```{{exec}}
1. Verify that the services are running:

   ```bash
   docker compose ps
   ```{{exec}}

   You should see something like:

   ```console
   docker compose ps
   NAME                            COMMAND                  SERVICE             STATUS              PORTS
   single-binary-alloy-1           "/bin/alloy run /etc…"   alloy               running             0.0.0.0:4319->4317/tcp, 0.0.0.0:12345->12345/tcp
   single-binary-grafana-1         "/run.sh"                grafana             running             0.0.0.0:3000->3000/tcp
   single-binary-k6-tracing-1      "/k6-tracing run /ex…"   k6-tracing          running
   single-binary-minio-1           "sh -euc 'mkdir -p /…"   minio               running             0.0.0.0:9001->9001/tcp
   single-binary-prometheus-1      "/bin/prometheus --c…"   prometheus          running             0.0.0.0:9090->9090/tcp
   single-binary-redpanda-1        "/entrypoint.sh redp…"   redpanda            running             0.0.0.0:9092->9092/tcp
   single-binary-redpanda-console-1 "/bin/sh -c 'echo \"$…" redpanda-console    running             0.0.0.0:8080->8080/tcp
   single-binary-tempo-1           "/tempo -target=all …"   tempo               running             0.0.0.0:3200->3200/tcp, 0.0.0.0:4317-4318->4317-4318/tcp
   single-binary-vulture-1         "/tempo-vulture -pro…"   vulture             running
   ```{{copy}}
