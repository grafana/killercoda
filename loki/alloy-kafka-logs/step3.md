# Step 3: Configure Alloy to ingest OpenTelemetry logs via Kafka

Next we will configure Alloy to also ingest OpenTelemetry logs via Kafka, we need to update the Alloy configuration file once again. We will add the new components to the `config.alloy`{{copy}} file along with the existing components.

**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**

## Source OpenTelemetry logs from Kafka

First, we will configure the OpenTelemetry Kafaka receiver. `otelcol.receiver.kafka`{{copy}} accepts telemetry data from a Kafka broker and forwards it to other `otelcol.*`{{copy}} components.

Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

```alloy
otelcol.receiver.kafka "default" {
  brokers          = ["kafka:9092"]
  protocol_version = "2.0.0"
  topic           = "otlp"
  encoding        = "otlp_proto"

  output {
    logs    = [otelcol.processor.batch.default.input]
  }
}
```{{copy}}

In this configuration:

- `brokers`{{copy}}: The Kafka brokers to connect to.

- `protocol_version`{{copy}}: The Kafka protocol version to use.

- `topic`{{copy}}: The Kafka topic to consume. In this case, we are consuming the `otlp`{{copy}} topic.

- `encoding`{{copy}}: The encoding of the incoming logs. Which decodes messages as OTLP protobuf.

- `output`{{copy}}: The list of receivers to forward the logs to. In this case, we are forwarding the logs to the `otelcol.processor.batch.default.input`{{copy}}.

For more information on the `otelcol.receiver.kafka`{{copy}} configuration, see the [OpenTelemetry Receiver Kafka documentation](https://grafana.com/docs/alloy/latest/reference/components/otelcol.receiver.kafka/).

## Batch OpenTelemetry logs before sending

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

## Write OpenTelemetry logs to Loki

Lastly, we will configure the OpenTelemetry exporter. `otelcol.exporter.otlphttp`{{copy}} accepts telemetry data from other otelcol components and writes them over the network using the OTLP HTTP protocol. We will use this exporter to send the logs to Loki’s native OTLP endpoint.

Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

```alloy
otelcol.exporter.otlphttp "default" {
  client {
    endpoint = "http://loki:3100/otlp"
  }
}
```{{copy}}

In this configuration:

- `client`{{copy}}: The client configuration for the exporter. In this case, we are sending the logs to the Loki OTLP endpoint.

For more information on the `otelcol.exporter.otlphttp`{{copy}} configuration, see the [OpenTelemetry Exporter OTLP HTTP documentation](https://grafana.com/docs/alloy/latest/reference/components/otelcol.exporter.otlphttp/).

## Reload the Alloy configuration to check the changes

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
