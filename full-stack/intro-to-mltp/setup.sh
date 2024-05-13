#!/usr/bin/env bash

echo "RUNNING SETUP SCRIPT"

# Clone the tutorial environment repository if it doesn't already exist
if [ ! -d "intro-to-mltp" ]; then
    git clone https://github.com/grafana/intro-to-mltp.git || { echo "Failed to clone repository"; exit 1; }
fi

# Enter the directory and switch to the required branch
cd intro-to-mltp && git checkout killercoda || { echo "Moving directory"; exit 1; }

echo "Building training instance...."
docker-compose -f docker-compose-no-beyla.yml up -d
echo "Catch any failed containers...."
docker-compose -f docker-compose-no-beyla.yml up -d


# Update and install required packages
echo "Updating and installing required packages..."
sudo apt-get update && sudo apt-get install -y figlet; clear; echo -e "\e[32m$(figlet -f standard 'Intro to')\e[0m"; echo -e "\e[33m$(figlet -f standard 'MLTP')\e[0m"