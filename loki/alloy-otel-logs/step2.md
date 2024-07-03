# Step 2: Configure Alloy to ingest OpenTelemetry logs

To configure Alloy to ingest OpenTelemetry logs, we need to update the Alloy configuration file. To start, we will update the `config.alloy`{{copy}} file to include the OpenTelemetry logs configuration.

{{< docs/ignore >}}
**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**
{{< /docs/ignore >}}

## Recive OpenTelemetry logs via gRPC and HTTP

First, we will configure the OpenTelemetry receiver. `otelcol.receiver.otlp`{{copy}} accepts logs in the OpenTelemetry format via HTTP and gRPC. We will use this receiver to receive logs from the Carnivorous Greenhouse application.

Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

```alloy
 otelcol.receiver.otlp "default" {
   http {}
   grpc {}

   output {
     logs    = [otelcol.processor.batch.default.input]
   }
 }
```{{copy}}

In this configuration:

- `http`{{copy}}: The HTTP configuration for the receiver. This configuration is used to receive logs in the OpenTelemetry format via HTTP.

- `grpc`{{copy}}: The gRPC configuration for the receiver. This configuration is used to receive logs in the OpenTelemetry format via gRPC.

- `output`{{copy}}: The list of processors to forward the logs to. In this case, we are forwarding the logs to the `otelcol.processor.batch.default.input`{{copy}}.

For more information on the `otelcol.receiver.otlp`{{copy}} configuration, see the [OpenTelemetry Receiver OTLP documentation](https://grafana.com/docs/alloy/latest/reference/components/otelcol.receiver.otlp/).

## Create batches of logs using a OpenTelemetry Processor

Next, we will configure a OpenTelemetry processor. `otelcol.processor.batch`{{copy}} accepts telemetry data from other otelcol components and places them into batches. Batching improves the compression of data and reduces the number of outgoing network requests required to transmit data. This processor supports both size and time based batching.

Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

```alloy
otelcol.processor.batch "default" {
    output {
        logs = [otelcol.exporter.otlphttp.default.input]
    }
}
```{{copy}}

In this configuration:

- `output`{{copy}}: The list of receivers to forward the logs to. In this case, we are forwarding the logs to the `otelcol.exporter.otlphttp.default.input`{{copy}}.

For more information on the `otelcol.processor.batch`{{copy}} configuration, see the [OpenTelemetry Processor Batch documentation](https://grafana.com/docs/alloy/latest/reference/components/otelcol.processor.batch/).

## Export logs to Loki using a OpenTelemetry Exporter

Lastly, we will configure the OpenTelemetry exporter. `otelcol.exporter.otlphttp`{{copy}} accepts telemetry data from other otelcol components and writes them over the network using the OTLP HTTP protocol. We will use this exporter to send the logs to Loki’s native OTLP endpoint.

Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

```alloy
otelcol.exporter.otlphttp "default" {
  client {
    endpoint = "http://loki:3100/otlp"
  }
}
```{{copy}}

For more information on the `otelcol.exporter.otlphttp`{{copy}} configuration, see the [OpenTelemetry Exporter OTLP HTTP documentation](https://grafana.com/docs/alloy/latest/reference/components/otelcol.exporter.otlphttp/).

## Reload the Alloy configuration

Once added, save the file. Then run the following command to request Alloy to reload the configuration:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

The new configuration will be loaded this can be verified by checking the Alloy UI: [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `config.alloy`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/config.alloy loki-fundamentals/config.alloy
curl -X POST http://localhost:12345/-/reload
```{{exec}}
