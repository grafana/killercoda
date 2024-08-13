# Set up a local Grafana instance

In this tutorial, you configure Alloy to collect logs from your local machine and send them to Loki.
You can use the following Docker Compose file to set up a local Grafana instance.
This Docker Compose file includes Loki and Prometheus configured as data sources.

> The interactive sandbox has a VSCode-like editor that allows you to access files and folders. To access this feature, click on the `Editor` tab. The editor also has a terminal that you can use to run commands. Since some commands assume you are within a specific directory, we recommend running the commands in `tab1`.
1. Create a new directory and save the Docker Compose file as `docker-compose.yml`{{copy}}.

   ```bash
   mkdir alloy-tutorial
   cd alloy-tutorial
   touch docker-compose.yml
   ```{{exec}}

1. Copy the following Docker Compose file into `docker-compose.yml`{{copy}}.
   > We recommend using the `Editor`{{copy}} tab to copy and paste the Docker Compose file. However, you can also use a terminal editor like `nano`{{copy}} or `vim`{{copy}}.

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
   ```{{copy}}

1. To start the local Grafana instance, run the following command.

   ```bash
    docker-compose up -d
   ```{{exec}}

1. Open [http://localhost:3000]({{TRAFFIC_HOST1_3000}}) in your browser to access the Grafana UI.
