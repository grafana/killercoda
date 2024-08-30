# Clone the Tempo repository and start Docker

1. Clone the Tempo repository:

   ```bash
   git clone https://github.com/grafana/tempo.git
   ```{{exec}}

1. Go into the examples directory:

   ```bash
   cd tempo/example/docker-compose/local
   ```{{exec}}

1. Start the services defined in the docker-compose file:

   ```bash
   docker compose up -d
   ```{{exec}}

1. Verify that the services are running using `docker compose ps`{{copy}}. You should see something like:

   ```console
   docker compose ps
   NAME                 COMMAND                  SERVICE             STATUS              PORTS
   local-grafana-1      "/run.sh"                grafana             running             0.0.0.0:3000->3000/tcp
   local-k6-tracing-1   "/k6-tracing run /ex…"   k6-tracing          running
   local-prometheus-1   "/bin/prometheus --c…"   prometheus          running             0.0.0.0:9090->9090/tcp
   local-tempo-1        "/tempo -config.file…"   tempo               running             0.0.0.0:3200->3200/tcp, 0.0.0.0:4317-4318->4317-4318/tcp, 0.0.0.0:9411->9411/tcp, 0.0.0.0:14268->14268/tcp
   ```{{copy}}
