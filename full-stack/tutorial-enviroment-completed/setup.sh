#!/usr/bin/env bash

echo "RUNNING SETUP SCRIPT"

# Clone the tutorial environment repository if it doesn't already exist
if [ ! -d "tutorial-environment" ]; then
    git clone https://github.com/grafana/tutorial-environment.git || { echo "Failed to clone repository"; exit 1; }
fi

# Enter the directory and switch to the required branch
cd tutorial-environment && git checkout killercoda || { echo "Failed to checkout branch"; exit 1; }

echo "Building training instance...."
docker-compose up -d || { echo "Failed to start docker containers"; exit 1; }

# Update and install required packages
echo "Updating and installing required packages..."
sudo apt-get update && sudo apt-get install -y toilet || { echo "Failed to install packages"; exit 1; }

# Display completion message
echo "SETUP COMPLETE"
clear
toilet -f smblock --filter metal "GRAFANA FUNDAMENTALS"
