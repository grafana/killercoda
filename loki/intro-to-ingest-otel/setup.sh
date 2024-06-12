#!/usr/bin/env bash

echo "RUNNING SETUP SCRIPT"

# Clone the tutorial environment repository if it doesn't already exist
if [ ! -d "loki-fundamentals" ]; then
    git clone https://github.com/grafana/loki-fundamentals.git || { echo "Failed to clone repository"; exit 1; }
fi

# Enter the directory and switch to the required branch
cd loki-fundamentals && git checkout intro-to-otel || { echo "Failed to checkout branch"; exit 1; }

echo "Building training instance...."

# Update and install required packages
echo "Updating and installing required packages..."
sudo apt-get update && apt install gpg -y; sudo mkdir -p /etc/apt/keyrings/; wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null; echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee /etc/apt/sources.list.d/grafana.list; sudo apt-get update && sudo apt-get install alloy; sudo apt-get install -y figlet; clear; echo -e "\e[32m$(figlet -f standard 'Intro to ingesting using')\e[0m"; echo -e "\e[33m$(figlet -f standard 'OpenTelemetry')\e[0m"