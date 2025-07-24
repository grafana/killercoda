#!/usr/bin/env bash

echo "--------------"
echo "4 steps setup:"
echo "--------------"

# LGTM
echo -e "\n>> Step 1: Setting up LGTM + Alloy stack..."
if docker compose -f /root/course/docker-compose.yaml up -d; then
    echo "LGTM ready"
    echo "You can start the workshop as step 1 is exploring the LGTM + Alloy stack."
else
    echo "Error: Failed to start LGTM stack"
    exit 1
fi

# Set up Java app
echo -e "\n>> Step 2: Setting up Java app..."
apt-get update -qq
apt-get install -y openjdk-17-jdk openjdk-17-jre > /dev/null 2>&1
echo "JDK installed"

if [ ! -d "/root/course/rolldice" ]; then
    echo "Error: /root/course/rolldice directory not found"
    exit 1
fi

cd /root/course/rolldice
chmod +x ./mvnw
chmod +x ./run.sh
if ./mvnw clean package -DskipTests > /dev/null 2>&1; then
    echo "Java app built successfully"
else
    echo "Error: Failed to build Java app"
    exit 1
fi

# Download OpenTelemetry agent
echo -e "\n>> Step 3: Downloading OpenTelemetry agent..."
version=v2.13.0
jar=opentelemetry-javaagent.jar
if curl -sL https://github.com/grafana/grafana-opentelemetry-java/releases/download/${version}/grafana-opentelemetry-java.jar -o ${jar}; then
    echo "OpenTelemetry agent downloaded"
else
    echo "Error: Failed to download OpenTelemetry agent"
fi

# Pull additional Docker images
echo -e "\n>> Step 4: Pulling k6 image for load testing..."
docker pull grafana/k6:latest
echo "Setup complete!"
