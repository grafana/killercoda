# Sending Logs to Loki via Kafka using Alloy

Alloy nativley supports receiving logs via Kafka. In this example, we will configure Alloy to recive logs via kafka using two different methods:

- [loki.source.kafka](https://grafana.com/docs/alloy/latest/reference/components/loki.source.kafka): reads messages from Kafka using a consumer group and forwards them to other `loki.*`{{copy}} components.

- [otelcol.receiver.kafka](https://grafana.com/docs/alloy/latest/reference/components/otelcol.receiver.kafka/): accepts telemetry data from a Kafka broker and forwards it to other `otelcol.*`{{copy}} components.

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

Each service generates logs that are sent to Alloy via Kafka. In this example, they are sent on two different topics:

- `loki`{{copy}}: This sends a structured log formatted message (json).

- `otlp`{{copy}}: This sends a serialized OpenTelemetry log message.

You would not typically do this within your own application, but for the purposes of this example we wanted to show how Alloy can handle different types of log messages over Kafka.
