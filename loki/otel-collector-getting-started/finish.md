# Summary

In this example, we configured the OpenTelemetry Collector to receive logs from an example application and send them to Loki using the native OTLP endpoint. Make sure to also consult the Loki configuration file `loki-config.yaml`{{copy}} to understand how we have configured Loki to receive logs from the OpenTelemetry Collector.

## Back to Docs

Head back to where you started from to continue with the Loki documentation: [Loki documentation](https://grafana.com/docs/loki/latest/send-data/otel)

# Further reading

For more information on the OpenTelemetry Collector and the native OTLP endpoint of Loki, refer to the following resources:

- [Loki OTLP endpoint](https://grafana.com/docs/loki/latest/send-data/otel/)

- [How is native OTLP endpoint different from Loki Exporter](https://grafana.com/docs/loki/latest/send-data/otel/native_otlp_vs_loki_exporter)

- [OpenTelemetry Collector Configuration](https://opentelemetry.io/docs/collector/configuration/)

# Complete metrics, logs, traces, and profiling example

If you would like to use a demo that includes Mimir, Loki, Tempo, and Grafana, you can use [Introduction to Metrics, Logs, Traces, and Profiling in Grafana](https://github.com/grafana/intro-to-mlt). `Intro-to-mltp`{{copy}} provides a self-contained environment for learning about Mimir, Loki, Tempo, and Grafana.

The project includes detailed explanations of each component and annotated configurations for a single-instance deployment. Data from `intro-to-mltp`{{copy}} can also be pushed to Grafana Cloud.
