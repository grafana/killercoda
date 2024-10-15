# Step 2: Configure Fluent Bit to send logs to Loki

To configure Fluent Bit to receive logs from our application, we need to provide a configuration file. This configuration file will define the components and their relationships. We will build the entire observability pipeline within this configuration file.

## Open your code editor and locate the `fluent-bit.conf`{{copy}} file

Fluent Bit requires a configuration file to define the components and their relationships. The configuration file is written using Fluent Bit configuration syntax. We will build the entire observability pipeline within this configuration file. To start, we will open the `fluent-bit.conf`{{copy}} file in the code editor:

> Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.

1. Expand the `loki-fundamentals`{{copy}} directory in the file explorer of the `Editor`{{copy}} tab.

1. Locate the `fluent-bit.conf`{{copy}} file in the top level directory, `loki-fundamentals`{{copy}}.

1. Click on the `fluent-bit.conf`{{copy}} file to open it in the code editor.

You will copy all of the configuration snippets into the `fluent-bit.conf`{{copy}} file.

## Receiving Fluent Bit protocal logs

The first step is to configure Fluent Bit to receive logs from the Carnivorous Greenhouse application. Since the application is instrumented with Fluent Bit logging framework, it will send logs using the forward protocol (unique to Fluent Bit). We will use the `forward`{{copy}} input plugin to receive logs from the application.

Now add the following configuration to the `fluent-bit.conf`{{copy}} file:

```conf
[INPUT]
    Name              forward
    Listen            0.0.0.0
    Port              24224
```{{copy}}

In this configuration:

- `Name`{{copy}}: The name of the input plugin. In this case, we are using the `forward`{{copy}} input plugin.

- `Listen`{{copy}}: The IP address to listen on. In this case, we are listening on all IP addresses.

- `Port`{{copy}}: The port to listen on. In this case, we are listening on port `24224`{{copy}}.

For more information on the `forward`{{copy}} input plugin, see the [Fluent Bit Forward documentation](https://docs.fluentbit.io/manual/pipeline/inputs/forward).

## Export logs to Loki using the official Loki output plugin

Lastly, we will configure Fluent Bit to export logs to Loki using the official Loki output plugin. The Loki output plugin allows you to send logs or events to a Loki service. It supports data enrichment with Kubernetes labels, custom label keys, and structured metadata.

Add the following configuration to the `fluent-bit.conf`{{copy}} file:

```conf
[OUTPUT]
    name   loki
    match  service.**
    host   loki
    port   3100
    labels agent=fluent-bit
    label_map_path /fluent-bit/etc/conf/logmap.json
```{{copy}}

In this configuration:

- `name`{{copy}}: The name of the output plugin. In this case, we are using the `loki`{{copy}} output plugin.

- `match`{{copy}}: The tag to match. In this case, we are matching all logs with the tag `service.**`{{copy}}.

- `host`{{copy}}: The hostname of the Loki service. In this case, we are using the hostname `loki`{{copy}}.

- `port`{{copy}}: The port of the Loki service. In this case, we are using port `3100`{{copy}}.

- `labels`{{copy}}: Additional labels to add to the logs. In this case, we are adding the label `agent=fluent-bit`{{copy}}.

- `label_map_path`{{copy}}: The path to the label map file. In this case, we are using the file `logmap.json`{{copy}}.

For more information on the `loki`{{copy}} output plugin, see the [Fluent Bit Loki documentation](https://docs.fluentbit.io/manual/pipeline/outputs/loki).

### `logmap.json`{{copy}} file

The `logmap.json`{{copy}} file is used to map the log fields to the Loki labels. In this tutorial we have pre-filled the `logmap.json`{{copy}} file with the following configuration:

```json
{
"service": "service_name",
"instance_id": "instance_id"
 }
```{{copy}}

This configuration maps the `service`{{copy}} field to the Loki label `service_name`{{copy}} and the `instance_id`{{copy}} field to the Loki label `instance_id`{{copy}}.

## Reload the Fluent Bit configuration

After adding the configuration to the `fluent-bit.conf`{{copy}} file, you will need to reload the Fluent Bit configuration. To reload the configuration, run the following command:

```bash
docker restart loki-fundamentals_fluent-bit_1
```{{exec}}

To verify that the configuration has been loaded successfully, you can check the Fluent Bit logs by running the following command:

```bash
docker logs loki-fundamentals_fluent-bit_1
```{{exec}}

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `config.alloy`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/fluent-bit.conf loki-fundamentals/fluent-bit.conf
docker restart loki-fundamentals_fluent-bit_1
```{{exec}}
