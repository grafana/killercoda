#!/usr/bin/env bash

# Function to show progress
show_progress() {
    echo "Background setup: $1"
}

# Pull additional Docker images
show_progress "Pulling k6 image for load testing..."
docker pull grafana/k6:latest

# Set up Java environment
show_progress "Setting up Java environment..."
apt-get update -qq
apt-get install -y openjdk-17-jdk openjdk-17-jre > /dev/null 2>&1

# Build and prepare demo application
show_progress "Building demo application..."
cd /root/course/rolldice
chmod +x ./mvnw
chmod +x ./run.sh
./mvnw clean package -DskipTests > /dev/null 2>&1

# Download OpenTelemetry agent
show_progress "Downloading OpenTelemetry agent..."
version=v2.13.0
jar=opentelemetry-javaagent.jar
curl -sL https://github.com/grafana/grafana-opentelemetry-java/releases/download/${version}/grafana-opentelemetry-java.jar -o ${jar}

show_progress "Setup complete! Demo environment is ready."
