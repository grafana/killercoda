# Step 2: Configuring the OpenTelemetry Collector

To configure the Collector to ingest OpenTelemetry logs from our application, we need to provide a configuration file. This configuration file will define the components and their relationships. We will build the entire observability pipeline within this configuration file.

## Open your Code Editor and Locate the `otel-config.yaml`{{copy}} file

The configuration file is written using yaml configuration syntax.To start, we will open the `otel-config.yaml`{{copy}} file in the code editor:

**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**

1. Expand the `loki-fundamentals`{{copy}} directory in the file explorer of the `Editor`{{copy}} tab.

1. Locate the `otel-config.yaml`{{copy}} file in the top level directory, `loki-fundamentals'.

1. Click on the `otel-config.yaml`{{copy}} file to open it in the code editor.

You will copy all three of the following configuration snippets into the `otel-config.yaml`{{copy}} file.

## Recive OpenTelemetry logs via gRPC and HTTP

First, we will configure the OpenTelemetry receiver. `otlp:`{{copy}} accepts logs in the OpenTelemetry format via HTTP and gRPC. We will use this receiver to receive logs from the Carnivorous Greenhouse application.

Now add the following configuration to the `otel-config.yaml`{{copy}} file:

```yaml
# Receivers
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318
```{{copy}}

In this configuration:

- `receivers`{{copy}}: The list of receivers to receive telemetry data. In this case, we are using the `otlp`{{copy}} receiver.

- `otlp`{{copy}}: The OpenTelemetry receiver that accepts logs in the OpenTelemetry format.

- `protocols`{{copy}}: The list of protocols that the receiver supports. In this case, we are using `grpc`{{copy}} and `http`{{copy}}.

- `grpc`{{copy}}: The gRPC protocol configuration. The receiver will accept logs via gRPC on `

- `http`{{copy}}: The HTTP protocol configuration. The receiver will accept logs via HTTP on `

- `endpoint`{{copy}}: The IP address and port number to listen on. In this case, we are listening on all IP addresses on port `4317`{{copy}} for gRPC and port `4318`{{copy}} for HTTP.

For more information on the `otlp`{{copy}} receiver configuration, see the [OpenTelemetry Receiver OTLP documentation](https://github.com/open-telemetry/opentelemetry-collector/blob/main/receiver/otlpreceiver/README.md).

## Create batches of logs using a OpenTelemetry Processor

Next, we will configure a OpenTelemetry processor. `batch:`{{copy}} accepts telemetry data from other `otelcol`{{copy}} components and places them into batches. Batching improves the compression of data and reduces the number of outgoing network requests required to transmit data. This processor supports both size and time based batching.

Now add the following configuration to the `otel-config.yaml`{{copy}} file:

```yaml
# Processors
processors:
  batch:
```{{copy}}

In this configuration:

- `processors`{{copy}}: The list of processors to process telemetry data. In this case, we are using the `batch`{{copy}} processor.

- `batch`{{copy}}: The OpenTelemetry processor that accepts telemetry data from other `otelcol`{{copy}} components and places them into batches.

For more information on the `batch`{{copy}} processor configuration, see the [OpenTelemetry Processor Batch documentation](https://github.com/open-telemetry/opentelemetry-collector/blob/main/processor/batchprocessor/README.md).

## Export logs to Loki using a OpenTelemetry Exporter

Lastly, we will configure the OpenTelemetry exporter. `otlphttp/logs:`{{copy}} accepts telemetry data from other `otelcol`{{copy}} components and writes them over the network using the OTLP HTTP protocol. We will use this exporter to send the logs to the Loki native OTLP endpoint.

Now add the following configuration to the `otel-config.yaml`{{copy}} file:

```yaml
# Exporters
exporters:
  otlphttp/logs:
    endpoint: "http://loki:3100/otlp"
    tls:
      insecure: true
```{{copy}}

In this configuration:

- `exporters`{{copy}}: The list of exporters to export telemetry data. In this case, we are using the `otlphttp/logs`{{copy}} exporter.

- `otlphttp/logs`{{copy}}: The OpenTelemetry exporter that accepts telemetry data from other `otelcol`{{copy}} components and writes them over the network using the OTLP HTTP protocol.

- `endpoint`{{copy}}: The URL to send the telemetry data to. In this case, we are sending the logs to the Loki native OTLP endpoint at `http://loki:3100/otlp`{{copy}}.

- `tls`{{copy}}: The TLS configuration for the exporter. In this case, we are setting `insecure`{{copy}} to `true`{{copy}} to disable TLS verification.

- `insecure`{{copy}}: Disables TLS verification. This is set to `true`{{copy}} as we are using an insecure connection.

For more information on the `otlphttp/logs`{{copy}} exporter configuration, see the [OpenTelemetry Exporter OTLP HTTP documentation](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter/README.md)

## Creating the Pipeline

Now that we have configured the receiver, processor, and exporter, we need to create a pipeline to connect these components. Add the following configuration to the `otel-config.yaml`{{copy}} file:

```yaml
# Pipelines
service:
  pipelines:
    logs:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlphttp/logs]
```{{copy}}

In this configuration:

- `pipelines`{{copy}}: The list of pipelines to connect the receiver, processor, and exporter. In this case, we are using the `logs`{{copy}} pipeline but there is also pipelines for metrics, traces, and continuous profiling.

- `receivers`{{copy}}: The list of receivers to receive telemetry data. In this case, we are using the `otlp`{{copy}} receiver component we created earlier.

- `processors`{{copy}}: The list of processors to process telemetry data. In this case, we are using the `batch`{{copy}} processor component we created earlier.

- `exporters`{{copy}}: The list of exporters to export telemetry data. In this case, we are using the `otlphttp/logs`{{copy}} component exporter we created earlier.

## Load the Configuration

Before you load the configuration, into the OpenTelemetry Collector compare your configuration with the completed configuration below:

```yaml
# Receivers
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318
        
# Processors
processors:
  batch:

# Exporters
exporters:
  otlphttp/logs:
    endpoint: "http://loki:3100/otlp"
    tls:
      insecure: true
      
# Pipelines
service:
  pipelines:
    logs:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlphttp/logs]
```{{copy}}

Next, we need apply the configuration to the OpenTelemetry Collector. To do this, we will restart the OpenTelemetry Collector container:

```bash
docker restart loki-fundamentals_otel-collector_1
```{{exec}}

This will restart the OpenTelemetry Collector container with the new configuration. You can check the logs of the OpenTelemetry Collector container to see if the configuration was loaded successfully:

```bash
docker logs loki-fundamentals_otel-collector_1
```{{exec}}

Within the logs, you should see the following message:

```console
2024-08-02T13:10:25.136Z        info    service@v0.106.1/service.go:225 Everything is ready. Begin running and processing data.
```{{exec}}

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `otel-config.yaml`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/otel-config.yaml loki-fundamentals/otel-config.yaml
docker restart loki-fundamentals_otel-collector_1
```{{exec}}
