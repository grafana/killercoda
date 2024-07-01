# Step 3: Start the Carnivorous Greenhouse

In this step, we will start the Carnivorous Greenhouse application. To start the application, run the following command:

**Note: This docker-compose file relies on the `loki-fundamentals_loki` docker network. If you have not started the observability stack, you will need to start it first.**


```bash
docker-compose -f loki-fundamentals/greenhouse/docker-compose-micro.yml up -d --build
```{{exec}}


This will start the following services:

```bash
 ✔ Container greenhouse-db-1                 Started                                                         
 ✔ Container greenhouse-websocket_service-1  Started 
 ✔ Container greenhouse-bug_service-1        Started
 ✔ Container greenhouse-user_service-1       Started
 ✔ Container greenhouse-plant_service-1      Started
 ✔ Container greenhouse-simulation_service-1 Started
 ✔ Container greenhouse-main_app-1           Started
```{{copy}}

Once started, you can access the Carnivorous Greenhouse application at [http://localhost:5005]({{TRAFFIC_HOST1_5005}}). Generate some logs by interacting with the application in the following ways:

- Create a user

- Log in

- Create a few plants to monitor

- Enable bug mode to activate the bug service. This will cause services to fail and generate additional logs.

Finally to view the logs in Loki, navigate to the Loki Logs Explore view in Grafana at [http://localhost:3000/a/grafana-lokiexplore-app/explore]({{TRAFFIC_HOST1_3000}}/a/grafana-lokiexplore-app/explore).
