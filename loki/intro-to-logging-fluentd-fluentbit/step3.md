# Step 3: Configure fluent Bit to send logs to Loki

In this step, we will configure Fluent Bit to send logs to Loki.

**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**

## Source Configuration

First, we need to configure Fluent Bit to recive logs from the Carnivorous Greenhouse application. Like Fluentd The carniverous greenhouse application is configured to send logs to Fluent Bit via the fluent python library. To recive these logs, we need to configure a source in Fluent bit to listen for logs on port `24224`{{copy}}. The type of source we will use is `forward`{{copy}}. Although we will use the `forward`{{copy}} input plugin once again the config is slightly different.

```apacheconf
[INPUT]
    Name              forward
    Listen            0.0.0.0
    Port              24224
```{{copy}}

Copy and paste the above configuration into the `fluent-bit.conf`{{copy}} file located at `loki-fundamentals/fluent-bit.conf`{{copy}}.

## Output Configuration

Next, we need to configure Fluent Bit to send logs to Loki. There are two ways to send logs to Loki using Fluent Bit. The first is to use the `loki`{{copy}} output plugin and the second is to use the `grafana-loki`{{copy}} output plugin. We recommend using the `grafana-loki`{{copy}} output plugin as it is more feature-rich and actively maintained by the community. The `grafana-loki`{{copy}} output plugin requires the following configuration:

```apacheconf
[OUTPUT]
    Name              grafana-loki
    Match             service.**
    Url               http://loki:3100/loki/api/v1/push
    Labels            {agent="fluent-bit"}
    LabelMapPath      /fluent-bit/etc/conf/logmap.json
```{{copy}}

Copy and paste the above configuration into the `fluent-bit.conf`{{copy}} file located at `loki-fundamentals/fluent-bit.conf`{{copy}}. This configuration has the following properties:

- `Name`{{copy}}: The name of the output plugin to use. In this case, we are using the `grafana-loki`{{copy}} output plugin.

- `Match`{{copy}}: The tag to match logs with. In this case, we are matching logs with the tag `service.**`{{copy}}.

- `Url`{{copy}}: The URL of the Loki instance to send logs to.

- `Labels`{{copy}}: Additional labels to add to the log stream.

- `LabelMapPath`{{copy}}: The path to the label map file. The label map file is used to map log fields to labels in Loki.

Lets quickly talk about LabelMapPath and logmap.json. The `LabelMapPath`{{copy}} is a path to a file that contains a mapping of log fields to labels in Loki. The `logmap.json`{{copy}} file is used to map log fields to labels in Loki. The `logmap.json`{{copy}} file should look like this:

```json
{
"service": "service_name",
"instance_id": "instance_id"
 }
```{{copy}}

The `logmap.json`{{copy}} file maps the `service`{{copy}} field in the log to the `service_name`{{copy}} label in Loki and the `instance_id`{{copy}} field in the log to the `instance_id`{{copy}} label in Loki.

# Restart the Fluentd Container

After configuring Fluentd, we need to restart the Fluentd container to apply the changes. To restart the Fluentd container, run the following command:

```bash
docker restart loki-fundamentals_fluent-bit_1
```{{exec}}

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `fluent-bit.conf`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/fluent-bit.conf loki-fundamentals/fluent-bit.conf
docker restart loki-fundamentals_fluent-bit_1
```{{exec}}
