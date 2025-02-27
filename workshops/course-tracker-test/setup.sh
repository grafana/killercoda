#!/bin/bash

# Define variables
COURSE="course-tracker-test"
VM_UUID=$(cat /sys/class/dmi/id/product_uuid)
BIN_DIR="/usr/local/bin"
SERVICE_NAME="course-monitor"
BINARY_NAME="alloy-linux-amd64"
CONFIG_NAME="config.alloy"
DOWNLOAD_URL="https://github.com/grafana/alloy/releases/download/v1.6.1/alloy-linux-amd64.zip"
CONFIG_URL="https://raw.githubusercontent.com/grafana/killercoda/refs/heads/staging/tools/course-tracker/config.alloy"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"

set -euf
# shellcheck disable=SC3040
(set -o pipefail 2> /dev/null) && set -o pipefail


sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

ARCH="$(dpkg --print-architecture)"
VERSION_CODENAME="$(source /etc/os-release && echo "${VERSION_CODENAME}")"
readonly ARCH VERSION_CODENAME

printf 'deb [arch=%s signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu %s stable' "${ARCH}" "${VERSION_CODENAME}" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install -y docker-compose-plugin

# Create a temporary directory for the download
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR" || exit 1

# Download and unzip Alloy
echo "Downloading Alloy..."
wget -q "$DOWNLOAD_URL" -O alloy.zip
unzip -q alloy.zip

# Move the binary to /usr/local/bin and make it executable
echo "Installing Alloy..."
sudo mv "$BINARY_NAME" "$BIN_DIR/$BINARY_NAME"
sudo chmod +x "$BIN_DIR/$BINARY_NAME"

# Download the configuration file
echo "Downloading configuration..."
sudo wget -q "$CONFIG_URL" -O "/etc/$CONFIG_NAME"

# Create the systemd service
echo "Creating systemd service..."
sudo bash -c "cat <<EOF > $SERVICE_FILE
[Unit]
Description=Course Monitor Service
After=network.target

[Service]
ExecStart=$BIN_DIR/$BINARY_NAME run /etc/$CONFIG_NAME
Restart=always
User=root
WorkingDirectory=$BIN_DIR
StandardOutput=journal
StandardError=journal
LimitNOFILE=65536
Environment=VM_UUID=$VM_UUID
Environment=COURSE=$COURSE
CUSTOM_ARGS="--server.http.listen-addr=0.0.0.0:12346"

[Install]
WantedBy=multi-user.target
EOF"


# Reload systemd, enable and start the service
echo "Enabling and starting the service..."
sudo systemctl daemon-reload
sudo systemctl enable "$SERVICE_NAME"
sudo systemctl start "$SERVICE_NAME"
export PROMPT_COMMAND='history -a'

echo "Service $SERVICE_NAME has been installed and started successfully." && clear && echo "Setup complete. You may now begin the tutorial."
