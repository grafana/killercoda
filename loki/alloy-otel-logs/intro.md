# Ingesting OpenTelemetry logs to Loki using Alloy

Alloy natively supports ingesting OpenTelemetry logs. In this example, we will configure Alloy to ingest OpenTelemetry logs to Loki.

## Dependencies

Before you begin, ensure you have the following to run the demo:

- Docker

- Docker Compose

## Scenario

In this scenario, we have a microservices application called the Carnivourse Greenhouse. This application consists of the following services:

- **User Service:** Mangages user data and authentication for the application. Such as creating users and logging in.

- **plant Service:** Manges the creation of new plants and updates other services when a new plant is created.

- **Simulation Service:** Generates sensor data for each plant.

- **Websocket Service:** Manages the websocket connections for the application.

- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.

- **Main App:** The main application that ties all the services together.

- **Database:** A database that stores user and plant data.

Each service generates logs using the OpenTelemetry SDK and exports to Alloy in the OpenTelemetry format. Alloy then ingests the logs and sends them to Loki. We will configure Alloy to ingest OpenTelemetry logs, send them to Loki, and view the logs in Grafana.
