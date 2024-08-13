# Getting started with the OpenTelemetry Collector and Loki tutorial

The OpenTelemetry Collector offers a vendor-agnostic implementation of how to receive, process and export telemetry data. With the introduction of the OTLP endpoint in Loki, you can now send logs from applications instrumented with OpenTelemetry to Loki using the OpenTelemetry Collector in native OTLP format.
In this example, we will teach you how to configure the OpenTelemetry Collector to receive logs in the OpenTelemetry format and send them to Loki using the OTLP HTTP protocol. This will involve configuring the following components in the OpenTelemetry Collector:

- **OpenTelemetry Receiver:** This component will receive logs in the OpenTelemetry format via HTTP and gRPC.

- **OpenTelemetry Processor:** This component will accept telemetry data from other `otelcol.*`{{copy}} components and place them into batches. Batching improves the compression of data and reduces the number of outgoing network requests required to transmit data.

- **OpenTelemetry Exporter:** This component will accept telemetry data from other `otelcol.*`{{copy}} components and write them over the network using the OTLP HTTP protocol. We will use this exporter to send the logs to the Loki native OTLP endpoint.

## Scenario

In this scenario, we have a microservices application called the Carnivourse Greenhouse. This application consists of the following services:

- **User Service:** Manages user data and authentication for the application. Such as creating users and logging in.

- **Plant Service:** Manages the creation of new plants and updates other services when a new plant is created.

- **Simulation Service:** Generates sensor data for each plant.

- **Websocket Service:** Manages the websocket connections for the application.

- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.

- **Main App:** The main application that ties all the services together.

- **Database:** A database that stores user and plant data.

Each service generates logs using the OpenTelemetry SDK and exports to the OpenTelemetry Collector in the OpenTelemetry format (otlp). The collector then ingests the logs and sends them to Loki.
