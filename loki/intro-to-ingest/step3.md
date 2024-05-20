
# Step 3: Adding Labels

Now that we have the logs being scraped and sent to Loki, we can start adding labels to the logs. This will allow us to filter and query the logs more effectively. We can do this by using the `loki.process` component.

## Loki Processor Component

The `loki.process` allows us to parse are log entries and add attributes such as labels to them. Here is an example of how we can use the `loki.process` component to extract the `level` and `logger` from the log entries and add them as labels:

```json
loki.process "add_new_label" {
    stage.logfmt {
        mapping = {
            "extracted_level" = "level",
            "extracted_logger" = "logger",
        }
    }

    stage.labels {
        values = {
            "level" = "extracted_level",
            "logger" = "extracted_logger",
        }
    }
     forward_to = [loki.write.local_loki.receiver]
}

```

Lets add this to our config:

```bash
echo 'loki.process "add_new_label" {
    stage.logfmt {
        mapping = {
            "extracted_level" = "level",
            "extracted_logger" = "logger",
        }
    }

    stage.labels {
        values = {
            "level" = "extracted_level",
            "logger" = "extracted_logger",
        }
    }
     forward_to = [loki.write.local_loki.receiver]
}' >> ./alloy-config.alloy
```{{exec}}

Reload Alloy with this config change:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

After reloading Alloy, we can see the new component in the Alloy UI:
[http://localhost:12345]({{TRAFFIC_HOST1_12345}})

## Replumbing the Components

Now that we have added the `loki.process` component, we need to replumb the components to make sure the logs are being processed correctly. We can do this by connecting the `loki.process` component to the `loki.source.file` component and making sure our `loki.write` component is receiving the processed logs.

```json
loki.source.file "local_files" {
    targets    = local.file_match.applogs.targets
    forward_to = [loki.process.add_new_label.receiver]
}


loki.process "add_new_label" {
    stage.logfmt {
        mapping = {
            "extracted_level" = "level",
            "extracted_logger" = "logger",
        }
    }

    stage.labels {
        values = {
            "level" = "extracted_level",
            "logger" = "extracted_logger",
        }
    }
     forward_to = [loki.write.local_loki.receiver]
}
```

Lets make these changes to our config:

```bash
sed -i -e '/loki.source.file "local_files" {/,/}/{s/loki\.write\.local_loki\.receiver/loki.process.add_new_label.receiver/}' ./alloy-config.alloy
```{{exec}}

Reload Alloy with this config change:

```bash
docker restart loki-fundamentals_alloy_1
```{{exec}}

**Note:** We are restarting the Alloy container to apply the changes. This due to the fact that docker does not recognize changes to the config file after sed has been run.

After reloading Alloy, we can see the new component in the Alloy UI:
[http://localhost:12345]({{TRAFFIC_HOST1_12345}})

## Querying the Logs

Now that we have added labels to our logs, we can query them more effectively. For example, we can query all the logs with the `level` set to `error` in Grafana:

[http://localhost:3000/explore]({{TRAFFIC_HOST1_3000}}/explore)
