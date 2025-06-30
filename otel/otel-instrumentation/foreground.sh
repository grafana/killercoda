#!/usr/bin/env bash

echo "Starting monitoring infrastructure..."

cd /root/course/lgtm
docker-compose -f docker-compose.yaml up -d

# Wait for critical services to be healthy
echo "Waiting for Grafana to be ready..."
until docker container inspect grafana 2>/dev/null | grep -q '"Status": "healthy"'; do
  printf "."
  sleep 2
done

echo ">> Monitoring infrastructure is ready!"
echo ">> Additional components are being set up in the background..."
