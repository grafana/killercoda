#!/usr/bin/env bash

echo "RUNNING SETUP SCRIPT"

# Clone the tutorial environment repository if it doesn't already exist
if [ ! -d "loki-fundamentals" ]; then
    git https://github.com/grafana/loki-fundamentals.git || { echo "Failed to clone repository"; exit 1; }
fi

# Enter the directory and switch to the required branch
cd loki-fundamentals && git checkout what-is-loki || { echo "Failed to checkout branch"; exit 1; }

echo "Building training instance...."
docker-compose up -d || { echo "Failed to start docker containers"; exit 1; }

# Update and install required packages
echo "Updating and installing required packages..."
sudo apt-get update && sudo apt-get install -y figlet; clear; echo -e "\e[32m$(figlet -f standard 'What is')\e[0m"; echo -e "\e[33m$(figlet -f standard 'Loki?')\e[0m"