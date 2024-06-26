# Set up a local Grafana instance

To allow {{< param “PRODUCT_NAME” >}} to write data to Loki running in the local Grafana stack, you can use the following Docker Compose file to set up a local Grafana instance alongside Loki and Prometheus, which are pre-configured as data sources.

```yaml
version: '3'
services:
  loki:
    image: grafana/loki:3.0.0
    ports:
      - "3100:3100"
    command: -config.file=/etc/loki/local-config.yaml
  prometheus:
    image: prom/prometheus:v2.47.0
    command:
      - --web.enable-remote-write-receiver
      - --config.file=/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"
  grafana:
    environment:
      - GF_PATHS_PROVISIONING=/etc/grafana/provisioning
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
    entrypoint:
      - sh
      - -euc
      - |
        mkdir -p /etc/grafana/provisioning/datasources
        cat <<EOF > /etc/grafana/provisioning/datasources/ds.yaml
        apiVersion: 1
        datasources:
        - name: Loki
          type: loki
          access: proxy
          orgId: 1
          url: http://loki:3100
          basicAuth: false
          isDefault: false
          version: 1
          editable: false
        - name: Prometheus
          type: prometheus
          orgId: 1
          url: http://prometheus:9090
          basicAuth: false
          isDefault: true
          version: 1
          editable: false
        EOF
        /run.sh
    image: grafana/grafana:11.0.0
    ports:
      - "3000:3000"
```

Run `docker compose up`{{copy}} to start your Docker container and open [http://localhost:3000]({{TRAFFIC_HOST1_3000}}) in your browser to view the Grafana UI.

{{< admonition type=“note” >}}
If you the following error when you start your Docker container, `docker: 'compose' is not a docker command`{{copy}}, use the command `docker-compose up`{{copy}} to start your Docker container.
{{< /admonition >}}
