# Sending logs to Loki using Fluent Bit tutorial

In this tutorial, you will learn how to send logs to Loki using Fluent Bit. Fluent Bit is a lightweight and fast log processor and forwarder that can collect, process, and deliver logs to various destinations. We will use the official Fluent Bit Loki output plugin to send logs to Loki.

## Scenario

In this scenario, we have a microservices application called the Carnivourse Greenhouse. This application consists of the following services:

- **User Service:** Manages user data and authentication for the application. Such as creating users and logging in.

- **Plant Service:** Manages the creation of new plants and updates other services when a new plant is created.

- **Simulation Service:** Generates sensor data for each plant.

- **Websocket Service:** Manages the websocket connections for the application.

- **Bug Service:** A service that when enabled, randomly causes services to fail and generate additional logs.

- **Main App:** The main application that ties all the services together.

- **Database:** A database that stores user and plant data.

Each service has been instrumented with the fluent bit logging framework to generate logs. If you would like to learn more about how the Carnivorous Greenhouse application was instrumented with Fluent Bit, refer to the [Carnivorous Greenhouse repository](https://github.com/grafana/loki-fundamentals/blob/fluentbit-official/greenhouse/loggingfw.py).
