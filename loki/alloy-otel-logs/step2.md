# Step 2: Configure Alloy to ingest OpenTelemetry logs

To configure Alloy to ingest OpenTelemetry logs, we need to update the Alloy configuration file. To start, we will update the `config.alloy`{{copy}} file to include the OpenTelemetry logs configuration.

## OpenTelelmetry Logs Receiver

First, we will configure the OpenTelemetry logs receiver. This receiver will accept logs via HTTP and gRPC.

1. Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:

   ```alloy
     otelcol.receiver.otlp "default" {
       http {}
       grpc {}

       output {
         logs    = [otelcol.processor.batch.default.input]
       }
     }
   ```{{copy}}

Once added, save the file. Then run the following command to request Alloy to reload the configuration:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

## OpenTelemetry Logs Processor

Next, we will configure the OpenTelemetry logs processor. This processor will batch the logs before sending them to the logs exporter.

1. Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:
   ```
   ```alloy
   ```

   otelcol.processor.batch “default” {
   output {
   logs = [otelcol.exporter.otlphttp.default.input]
   }
   }
   ```
   <!-- Killercoda copy END -->

   ```{{copy}}

Once added, save the file. Then run the following command to request Alloy to reload the configuration:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

## OpenTelemetry Logs Exporter

Lastly, we will configure the OpenTelemetry logs exporter. This exporter will send the logs to Loki.

1. Open the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory and copy the following configuration:
   ```
   ```alloy
   otelcol.exporter.otlphttp "default" {
     client {
       endpoint = "http://loki:3100/otlp"
     }
   }
   ```
   ```

Once added, save the file. Then run the following command to request Alloy to reload the configuration:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `config.alloy`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/config.alloy loki-fundamentals/config.alloy
```{{exec}}