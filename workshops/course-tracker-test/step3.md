# Collect logs from a sample application

Currently, the Loki stack is collecting logs about itself. To provide a more realistic example, you can deploy a sample application that generates logs. The sample application is called **The Carnivourous Greenhouse**, a microservices application that allows users to login and simulate a greenhouse with carnivorous plants to monitor. The application consists of seven services:

- **User Service:** Manages user data and authentication for the application. Such as creating users and logging in.

- **Plant Service:** Manages the creation of new plants and updates other services when a new plant is created.

- **Simulation Service:** Generates sensor data for each plant.

- **WebSocket Service:** Manages the websocket connections for the application.

- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.

- **Main App:** The main application that ties all the services together.

- **Database:** A PostgreSQL database that stores user and plant data.

The architecture of the application is shown below:

![Sample Microservice Architecture](https://grafana.com/media/docs/loki/get-started-architecture.png)

To deploy the sample application, follow these steps:

1. With `loki-fundamentals`{{copy}} as the current working directory, deploy the sample application using Docker Compose:

   ```bash
   docker compose -f greenhouse/docker-compose-micro.yml up -d --build  
   ```{{exec}}

   > **Note:**
   > This may take a few minutes to complete since the images for the sample application need to be built. Go grab a coffee and come back.

   Once the command completes, you should see a similar output:

   ```console
     ✔ bug_service                                Built     0.0s 
     ✔ main_app                                   Built     0.0s 
     ✔ plant_service                              Built     0.0s 
     ✔ simulation_service                         Built     0.0s 
     ✔ user_service                               Built     0.0s 
     ✔ websocket_service                          Built     0.0s 
     ✔ Container greenhouse-websocket_service-1   Started   0.7s 
     ✔ Container greenhouse-db-1                  Started   0.7s 
     ✔ Container greenhouse-user_service-1        Started   0.8s 
     ✔ Container greenhouse-bug_service-1         Started   0.8s 
     ✔ Container greenhouse-plant_service-1       Started   0.8s 
     ✔ Container greenhouse-simulation_service-1  Started   0.7s 
     ✔ Container greenhouse-main_app-1            Started   0.7s
   ```{{copy}}

1. To verify the sample application is running, open a browser and navigate to [http://localhost:5005]({{TRAFFIC_HOST1_5005}}). You should see the login page for the Carnivorous Greenhouse application.

   ![Greenhouse Home Page](https://grafana.com/media/docs/loki/get-started-login.png)

   Now that the sample application is running, run some actions in the application to generate logs. Here is a list of actions:

   1. **Create a user:** Click **Sign Up** and create a new user. Add a username and password and click **Sign Up**.

   1. **Login:** Use the username and password you created to login. Add the username and password and click **Login**.

   1. **Create a plant:** Once logged in, give your plant a name, select a plant type and click **Add Plant**. Do this a few times if you like.

Your greenhouse should look something like this:

![Greenhouse Dashboard](https://grafana.com/media/docs/loki/get-started-greenhouse.png)

Now that you have generated some logs, you can return to the Grafana Logs Drilldown page [http://localhost:3000/a/grafana-lokiexplore-app]({{TRAFFIC_HOST1_3000}}/a/grafana-lokiexplore-app). You should see seven new services such as `greenhouse-main_app-1`{{copy}}, `greenhouse-plant_service-1`{{copy}}, `greenhouse-user_service-1`{{copy}}, etc.
