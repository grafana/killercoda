> We strongly encourage you to **open 2 terminal tabs**. You will use one to run the application and another to execute other commands in the tutorial.

# Install and instrument an application

We'll be working with "Rolldice", a simple Java application that we'll instrument with OpenTelemetry.

> Note: The Rolldice app is sourced from the [OpenTelemetry Workshop repository by Grafana Labs](https://github.com/grafana/opentelemetry-workshop).

## Setup Steps

1. Launch the application:
   ```bash
   cd ~/course/rolldice/
   ./run.sh
   ```{{exec}}

   Wait until you see the Tomcat server startup message. Then, in your second terminal tab, test the application:

   ```bash
   curl localhost:8080/rolldice
   ```{{exec}}

2. Stop the application in your first terminal tab (`Ctrl + C`)

The application is a simple server that returns a random number between 1 and 6 when requested.

## Zero-code Instrumentation

One of OpenTelemetry's powerful features is its ability to instrument Java applications automatically without code changes.

### 1. Configure Environment Variables

Set up the required OpenTelemetry configuration in a terminal:

```bash
export NAMESPACE="opentelemetry-test-learning"
export OTEL_RESOURCE_ATTRIBUTES="service.name=rolldice,deployment.environment=lab,service.namespace=${NAMESPACE},service.version=0.0.1,service.instance.id=${HOSTNAME}:8080"
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
export OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:4317"
```{{exec}}

This will create environment variables that our OpenTelemetry runtime will read and reuse.

What's happening here? We are configuring the OpenTelemetry Java agent to attach these OpenTelemetry _resource attributes_ to our signals:

| Attribute | Value | Description |
|-----------|-------|-------------|
| service.name | rolldice | The canonical name of our application |
| deployment.environment | lab | Environment where the app runs (e.g., "production", "test", "development") |
| service.instance.id | (your IDE's hostname) | Uniquely identifies this instance, useful for multi-instance deployments |
| service.namespace | opentelemetry-test-learning | Groups related applications in the same environment |

> OpenTelemetry components **often use environment variables** for configuration. The default value for `OTEL_EXPORTER_OTLP_ENDPOINT` assumes that you want to send telemetry to an OpenTelemetry collector on `localhost`. We could omit this environment variable entirely, but we're including it explicitly here to make it clear what's happening. 
In production, you might set this value to `http://alloy.mycompany.com:4317`, or wherever your Alloy instance is located.

### 2. Enable the OpenTelemetry Agent

Modify `run.sh` (located at `~/course/rolldice/run.sh`) to include the OpenTelemetry Java agent. Change the last line to: `java -javaagent:opentelemetry-javaagent.jar -jar ./target/rolldice-0.0.1-SNAPSHOT.jar`

The `-javaagent` argument tells the Java process to attach an agent when the program starts. Agents are other Java programs that can interact with and inspect the program that's running.

1. If you didn't stop the application, stop it now. (`Ctrl + C` in your terminal tab running the app)
2. Start it again with `./run.sh` in your 1st terminal tab

### 3. Generate Sample Data using k6

k6 is a load-testing tool. We provide a k6 script to execute.

In another terminal, run the k6 load test using Docker:
```bash
docker run --rm -i --network=host grafana/k6:latest run - < ~/course/load-test.js
```{{exec}}

This command:
- Launches k6 in Docker
- Connects to the host network to access localhost
- Executes the test script for 5 minutes
- Automatically cleans up the container after completion

## View the Results

1. Open your [Grafana instance]({{TRAFFIC_HOST1_3000}})

2. From the main menu, go to **Explore**.
   - In Drilldown > Metrics: search for `jvm`. You will see some statistics about the runtime.
   - In Drilldown > Logs: view log lines generated when a dice is rolled
   - In Drilldown > Traces: see all requests made to the _Rolldice_ application

> It can take some time for the signals to come in. You can see new metrics with no value yet.
