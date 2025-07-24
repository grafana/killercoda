# Install Grafana Alloy and configure OpenTelemetry pipelines

First, we need an OpenTelemetry Collector. In this tutorial, we will use Grafana Alloy, an alternative distribution from Grafana. It has been started with your stack.

Check Alloy's current configuration [on port 12345]({{TRAFFIC_HOST1_12345}}). Go to the Graph tab to better understand what is happening.

## Explanations

Our Alloy has a single and simple pipeline:
- Alloy starts a Prometheus exporter to expose metrics (about itself in this case)
- A discovery component is used to find where the exporter is located. It also adds some metadata to the metrics when scraped
- A scrape job to periodically get the metrics from the targets listed in the discovery component
- A relabel job to filter out some metrics that we don't want (more specifically, to keep only the ones we want in this case)
- A remote-write component to send those metrics to Mimir

Alloy will be the gateway for all signals and their processing. Let's configure it to accept metrics, logs, and traces through **OpenTelemetry Protocol (OTLP)** and push them to our backends.

## Configure our OpenTelemetry pipelines

### Configuration file

Let's modify our config file at `~/course/config.alloy`

> Killercoda includes an IDE interface. It's the first tab in the terminal window.

We will replace the current configuration with another one that has OpenTelemetry support. Open `~/course/opentelemetry.alloy` to explore it.

```bash
mv ~/course/opentelemetry.alloy ~/course/config.alloy
```{{exec}}

### Applying changes / Restart Alloy

Restart Alloy to apply your changes:

```bash
docker restart alloy
```{{exec}}

Check that your pipelines are up in [Alloy]({{TRAFFIC_HOST1_12345}})

## Understanding the pipelines

This graph is more complex than before. Here's what's happening:

1. **OTLP Receiver**: Listens for OTLP messages (logs, metrics, and traces)
2. **Resource Detection**: Enriches data with host metadata (e.g., hostname) that the app might not be aware of. From here, each signal follows a different path:
   - **Metrics Pipeline**:
     - Processor: to add metadata
     - Then sent to batch processor
   - **Logs Pipeline**: 
     - Directly sent to batch processor
   - **Traces Pipeline**: Split into two streams:
     - Sent to batch processor for backend storage
     - Sent to host info to generate usage metrics

3. **Batch Processor**: 
   - Groups messages to optimize network usage
   - Send signals to the LGTM stack

4. **Backend Communication**:
   - OTLPHTTP: Sends data to OpenTelemetry-compatible APIs

## Conclusion

Alloy is now configured to:
- Receive telemetry data via OTLP protocol on:
  - Port 4317 (gRPC)
  - Port 4318 (HTTP/protobuf)
- Enrich the data with host metadata
- Forward processed data to the LGTM stack

> The app could technically send telemetry directly to the LGTM stack. We use Alloy because:
> - It's a good way to visualize how a pipeline looks like
> - In prod, we recommend a central Collector (or cluster of Collectors) from which you can control data without the need to modify all apps and infra in your org

Let's proceed to instrument our application.
