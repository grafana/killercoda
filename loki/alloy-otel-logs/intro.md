# Sending OpenTelemetry logs to Loki using Alloy

Alloy natively supports receiving logs in the OpenTelemetry format. This allows you to send logs from applications instrumented with OpenTelemetry to Alloy, which can then be sent to Loki for storage and visualization in Grafana. In this example, we will make use of 3 Alloy components to achieve this:

- **OpenTelemetry Receiver:** This component will receive logs in the OpenTelemetry format via HTTP and gRPC.

- **OpenTelemetry Processor:** This component will accept telemetry data from other `otelcol.*`{{copy}} components and place them into batches. Batching improves the compression of data and reduces the number of outgoing network requests required to transmit data.

- **OpenTelemetry Exporter:** This component will accept telemetry data from other `otelcol.*`{{copy}} components and write them over the network using the OTLP HTTP protocol. We will use this exporter to send the logs to Loki’s native OTLP endpoint.

## Scenario

In this scenario, we have a microservices application called the Carnivourse Greenhouse. This application consists of the following services:

- **User Service:** Manages user data and authentication for the application. Such as creating users and logging in.

- **Plant Service:** Manages the creation of new plants and updates other services when a new plant is created.

- **Simulation Service:** Generates sensor data for each plant.

- **Websocket Service:** Manages the websocket connections for the application.

- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.

- **Main App:** The main application that ties all the services together.

- **Database:** A database that stores user and plant data.

Each service generates logs using the OpenTelemetry SDK and exports to Alloy in the OpenTelemetry format. Alloy then ingests the logs and sends them to Loki. We will configure Alloy to ingest OpenTelemetry logs, send them to Loki, and view the logs in Grafana.
