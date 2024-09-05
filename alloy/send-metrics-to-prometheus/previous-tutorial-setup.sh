#!/bin/bash

# Define the content of the docker-compose.yml file
compose_content=$(cat <<EOF
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
        cat <<EOS > /etc/grafana/provisioning/datasources/ds.yaml
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
        EOS
        /run.sh
    image: grafana/grafana:11.0.0
    ports:
      - "3000:3000"
EOF
)

# Define the content of the config.alloy file
alloy_content=$(cat <<EOF
local.file_match "local_files" {
    path_targets = [{"__path__" = "/var/log/*.log"}]
    sync_period = "5s"
}

loki.source.file "log_scrape" {
    targets    = local.file_match.local_files.targets
    forward_to = [loki.process.filter_logs.receiver]
    tail_from_end = true
}

loki.process "filter_logs" {
    stage.drop {
        source = ""
        expression  = ".*Connection closed by authenticating user root"
        drop_counter_reason = "noisy"
    }
    forward_to = [loki.write.grafana_loki.receiver]
}

loki.write "grafana_loki" {
    endpoint {
      url = "http://localhost:3100/loki/api/v1/push"
    }
}
EOF
)

# Create the docker-compose.yml file and add the content to it
echo "$compose_content" > docker-compose.yml
echo "docker-compose.yml has been created."

# Create the config.alloy file and add the content to it
echo "$alloy_content" > config.alloy
echo "config.alloy has been created."

# Install Alloy
sudo apt install gpg -y && \
sudo mkdir -p /etc/apt/keyrings/ && \
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null && \
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee /etc/apt/sources.list.d/grafana.list && \
sudo apt-get update && \
sudo apt-get install alloy -y && \
# Modify the Alloy service configuration to listen on the desired port
sudo sed -i -e 's/CUSTOM_ARGS=""/CUSTOM_ARGS="--server.http.listen-addr=0.0.0.0:12345"/' /etc/default/alloy && \
# Enable and start the Alloy service
sudo systemctl enable alloy && \
sudo systemctl start alloy.service && \
clear && \
echo "Alloy has been installed. You may now start the tutorial."
