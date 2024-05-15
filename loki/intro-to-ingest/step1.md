
# Step 1: Demo environment setup

The first thing we are going to do manually is spin up our Docker environment. This environment will be used to demonstrate how to ingest logs into Loki using the new Grafana Labs Alloy collector.

## Docker Compose

We will be using Docker Compose to spin up our environment. Run the following command to start the environment:

```bash
docker-compose up -d
```{{exec}}
This will build the Docker images and start the containers. Once the containers are up and running. To check the status of the containers, run the following command:

```bash
docker ps
```{{exec}}

Which should output something similar to:

```plaintext
CONTAINER ID   IMAGE                                 ...   STATUS         
392f00a18cf4   grafana/alloy:latest                  ...   Up 34 seconds 
60f6abe649a5   loki-fundamentals_carnivorous-garden  ...   Up 36 seconds  
c4a9ca220b0f   grafana/loki:main-e9b6ce9             ...   Up 36 seconds   
a32e179a44af   grafana/grafana:11.0.0                ...   Up 35 seconds 
```

## Access our applications

There are three application UI's that we will be using in this demo: Grafana, Alloy, and Carnivorous Greenhouse. Lets check if they are up and running:
* Grafana: [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
* Alloy: [http://localhost:12345]({{TRAFFIC_HOST1_12345}})
* Carnivorous Greenhouse: [http://localhost:5005]({{TRAFFIC_HOST1_5005}})



