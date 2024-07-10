# Step 2: Configure fluentd to send logs to Loki

In this step, we will configure Fluentd to send logs to Loki.

**Note: Killercoda has an inbuilt Code editor which can be accessed via the `Editor`{{copy}} tab.**

## Source Configuration

First, we need to configure Fluentd to recive logs from the Carnivorous Greenhouse application. The carniverous greenhouse application is configured to send logs to Fluentd via the fluent python library. To recive these logs, we need to configure a source in Fluentd to listen for logs on port `24224`{{copy}}. The type of source we will use is `forward`{{copy}}.

```apacheconf
<source>
  @type forward
  port 24224
  bind 0.0.0.0
</source>
```{{copy}}

Copy and paste the above configuration into the `fluentd.conf`{{copy}} file located at `loki-fundamentals/fluentd.conf`{{copy}}.

## Output Configuration

Next, we need to configure Fluentd to send logs to Loki. We will use the `loki`{{copy}} output plugin to send logs to Loki. The `loki`{{copy}} output plugin requires the following configuration:

```apacheconf
<match service.**>
  @type loki
  url "http://loki:3100"
  extra_labels {"agent": "fluentd"}
  <label>
   service_name $.service
   instance_id $.instance_id
  </label>
  <buffer>
    flush_interval 10s
    flush_at_shutdown true
    chunk_limit_size 1m
  </buffer>
</match>
```{{copy}}

Copy and paste the above configuration into the `fluentd.conf`{{copy}} file located at `loki-fundamentals/fluentd.conf`{{copy}}. This configuration has the following properties:

- `@type`{{copy}}: The type of output plugin to use. In this case, we are using the `loki`{{copy}} output plugin.

- `url`{{copy}}: The URL of the Loki instance to send logs to.

- `extra_labels`{{copy}}: Additional labels to add to the log stream.

- `label`{{copy}}: The labels to use for the log stream.

- `buffer`{{copy}}: The buffer configuration for the output plugin.
  Note that the `service.**`{{copy}} in the match directive is a tag that will match all logs with the tag `service`{{copy}}. We set the tag in the Carnivorous Greenhouse application to `service.<service_name>`{{copy}}.

# Restart the Fluentd Container

After configuring Fluentd, we need to restart the Fluentd container to apply the changes. To restart the Fluentd container, run the following command:

```bash
docker restart loki-fundamentals-fluentd-1
```{{exec}}

# Stuck? Need help?

If you get stuck or need help creating the configuration, you can copy and replace the entire `fluentd.conf`{{copy}} using the completed configuration file:

```bash
cp loki-fundamentals/completed/fluentd.conf loki-fundamentals/fluentd.conf
docker restart loki-fundamentals-fluentd-1
```{{exec}}
