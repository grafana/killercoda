#!/usr/bin/env bash

echo "RUNNING SETUP SCRIPT"

# Clone the tutorial environment repository if it doesn't already exist
if [ ! -d "tutorial-environment" ]; then
    git clone https://github.com/grafana/intro-to-mltp.git || { echo "Failed to clone repository"; exit 1; }
fi

# Enter the directory and switch to the required branch
cd intro-to-mltp && git checkout killercoda || { echo "Moving directory"; exit 1; }

echo "Building training instance...."
docker-compose -f docker-compose-no-beyla.yml up -d
