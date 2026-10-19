# Clone the Tempo repository and start Docker

This quick start guide uses the `single-binary`{{copy}} docker-compose example. Check out our [collection of examples](https://github.com/grafana/tempo/blob/main/example/docker-compose) to explore various
operating modes and configurations.

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
   NAME                         IMAGE                                                                                               COMMAND                  SERVICE      CREATED         STATUS         PORTS
   single-binary-alloy-1        docker.io/grafana/alloy@sha256:4f6ddc56ffdcf8a6316748fc5162972e20cb301523cac1bb4a31957df733ae9b     "run /etc/alloy/conf…"   alloy        7 seconds ago   Up 6 seconds   0.0.0.0:12345->12345/tcp, 0.0.0.0:4319->4317/tcp
   single-binary-grafana-1      docker.io/grafana/grafana@sha256:121a7a9ece6dc10b969f1f96eed64b4f07dfac0d0b8abc070f7cb83bbde86f63   ""                       grafana      7 seconds ago   Up 6 seconds   0.0.0.0:3000->3000/tcp
   single-binary-k6-tracing-1   ghcr.io/grafana/xk6-client-tracing:v0.0.9                                                           "run /example-script…"   k6-tracing   7 hours ago     Up 6 seconds
   single-binary-prometheus-1   docker.io/prom/prometheus:latest                                                                    "--config.file=/etc/…"   prometheus   7 seconds ago   Up 6 seconds   0.0.0.0:9090->9090/tcp
   single-binary-tempo-1        docker.io/grafana/tempo:3.0.0                                                                       "-target=all -config…"   tempo        7 seconds ago   Up 6 seconds   0.0.0.0:3200->3200/tcp, 0.0.0.0:4317-4318->4317-4318/tcp
   single-binary-vulture-1      docker.io/grafana/tempo-vulture:3.0.0                                                               "-prometheus-listen-…"   vulture      7 hours ago     Up 6 seconds
   ```{{copy}}
