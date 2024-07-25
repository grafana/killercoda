# Step 2: Configure Alloy to ingest raw Kafka logs

In this first step, we will configure Alloy to ingest raw Kafka logs. To do this, we will update the `config.alloy`{{copy}} file to include the Kafka logs configuration.

## Open your Code Editor and Locate the `config.alloy`{{copy}} file

Grafana Alloy requires a configuration file to define the components and their relationships. The configuration file is written using Alloy configuration syntax. We will build the entire observability pipeline within this configuration file. To start, we will open the `config.alloy`{{copy}} file in the code editor:

**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**

1. Expand the `loki-fundamentals`{{copy}} directory in the file explorer of the `Editor`{{copy}} tab.

1. Locate the `config.alloy`{{copy}} file in the `loki-fundamentals`{{copy}} directory (Top level directory).

1. Click on the `config.alloy`{{copy}} file to open it in the code editor.

You will copy all three of the following configuration snippets into the `config.alloy`{{copy}} file.

## Source logs from kafka

First, we will configure the Loki Kafka source. `loki.source.kafka`{{copy}} reads messages from Kafka using a consumer group and forwards them to other `loki.*`{{copy}} components.

The component starts a new Kafka consumer group for the given arguments and fans out incoming entries to the list of receivers in `forward_to`{{copy}}.

Add the following configuration to the `config.alloy`{{copy}} file:

```alloy
loki.source.kafka "raw" {
  brokers                = ["kafka:9092"]
  topics                 = ["loki"]
  forward_to             = [loki.write.http.receiver]
  relabel_rules          = loki.relabel.kafka.rules
  version                = "2.0.0"
  labels                = {service_name = "raw_kafka"}
}
```{{copy}}

In this configuration:

- `brokers`{{copy}}: The Kafka brokers to connect to.

- `topics`{{copy}}: The Kafka topics to consume. In this case, we are consuming the `loki`{{copy}} topic.

- `forward_to`{{copy}}: The list of receivers to forward the logs to. In this case, we are forwarding the logs to the `loki.write.http.receiver`{{copy}}.

- `relabel_rules`{{copy}}: The relabel rules to apply to the incoming logs. This can be used to generate labels from the temporary internal labels that are added by the Kafka source.

- `version`{{copy}}: The Kafka protocol version to use.

- `labels`{{copy}}: The labels to add to the incoming logs. In this case, we are adding a `service_name`{{copy}} label with the value `raw_kafka`{{copy}}. This will be used to identify the logs from the raw Kafka source in the Log Explorer App in Grafana.

For more information on the `loki.source.kafka`{{copy}} configuration, see the [Loki Kafka Source documentation](https://grafana.com/docs/alloy/latest/reference/components/loki.source.kafka/).

## Create a dynamic relabel based on Kafka topic

Next, we will configure the Loki relabel rules. The `loki.relabel`{{copy}} component rewrites the label set of each log entry passed to its receiver by applying one or more relabeling rules and forwards the results to the list of receivers in the component’s arguments. In our case we are directly calling the rule from the `loki.source.kafka`{{copy}} component.

Now add the following configuration to the `config.alloy`{{copy}} file:

```alloy
loki.relabel "kafka" {
  forward_to      = [loki.write.http.receiver]
  rule {
    source_labels = ["__meta_kafka_topic"]
    target_label  = "topic"
  }
}
```{{copy}}

In this configuration:

- `forward_to`{{copy}}: The list of receivers to forward the logs to. In this case, we are forwarding the logs to the `loki.write.http.receiver`{{copy}}. Though in this case, we are directly calling the rule from the `loki.source.kafka`{{copy}} component. So `forward_to`{{copy}} is being used as a placeholder as it is required by the `loki.relabel`{{copy}} component.

- `rule`{{copy}}: The relabeling rule to apply to the incoming logs. In this case, we are renaming the `__meta_kafka_topic`{{copy}} label to `topic`{{copy}}.

For more information on the `loki.relabel`{{copy}} configuration, see the [Loki Relabel documentation](https://grafana.com/docs/alloy/latest/reference/components/loki.relabel/).

## Write logs to Loki

Lastly, we will configure the Loki write component. `loki.write`{{copy}} receives log entries from other loki components and sends them over the network using the Loki logproto format.

And finally, add the following configuration to the `config.alloy`{{copy}} file:

```alloy
loki.write "http" {
  endpoint {
    url = "http://loki:3100/loki/api/v1/push"
  }
}
```{{copy}}

In this configuration:

- `endpoint`{{copy}}: The endpoint to send the logs to. In this case, we are sending the logs to the Loki HTTP endpoint.

For more information on the `loki.write`{{copy}} configuration, see the [Loki Write documentation](https://grafana.com/docs/alloy/latest/reference/components/loki.write/).

## Reload the Alloy configuration to check the changes

Once added, save the file. Then run the following command to request Alloy to reload the configuration:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

The new configuration will be loaded.  You can verify this by checking the Alloy UI: [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `config.alloy`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/config-raw.alloy loki-fundamentals/config.alloy
curl -X POST http://localhost:12345/-/reload
```{{exec}}
