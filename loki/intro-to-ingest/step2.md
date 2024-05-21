
# Step 2: Building out the Grafana Alloy config

We are going to start by building out the Grafana Alloy config. This config will be used to collect logs from our Carnivorous Greenhouse application. We will build this out step by step, due to the flexibility of Alloy and Loki, we can iterate through this process as we go and see the changes in real-time.

## Directory Discovery

The first thing we need to do is to discover the directories that contain the logs we want to collect. We can do this by using the `local.file_match`.

```json
local.file_match "applogs" {
    path_targets = [{"__path__" = "/tmp/app-logs/app.log"}]
    sync_period = "5s"
}
```
Lets add this to our `alloy.yaml` file.

```bash
echo 'local.file_match "applogs" {
    path_targets = [{"__path__" = "/tmp/app-logs/app.log"}]
    sync_period = "5s"
}' >> ./alloy-config.alloy
```{{exec}}

We can now reload Alloy with this config.

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

After reloading Alloy, we can see the new component in the Alloy UI:
[http://localhost:12345]({{TRAFFIC_HOST1_12345}})

## Writing to Loki

Next we we are going to establish a connection to Loki and write the logs to it. We can do this by using the `loki_push` component.

```json
loki.write "local_loki" {
    endpoint {
        url = "http://loki:3100/loki/api/v1/push"
    }
}
```
Lets add this to our config:
```bash
echo 'loki.write "local_loki" {
    endpoint {
        url = "http://loki:3100/loki/api/v1/push"
    }
}' >> ./alloy-config.alloy
```{{exec}}

Reload Alloy with this config change:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

After reloading Alloy, we can see the new component in the Alloy UI:
[http://localhost:12345]({{TRAFFIC_HOST1_12345}})


## Scraping the logs

Now that we have established a connection to Loki, we can start scraping the logs. We can do this by using the `loki.source.file` component.

```json
loki.source.file "local_files" {
    targets    = local.file_match.applogs.targets
    forward_to = [loki.write.local_loki.receiver]
}
```

Lets add this to our config:
```bash
echo 'loki.source.file "local_files" {
    targets    = local.file_match.applogs.targets
    forward_to = [loki.write.local_loki.receiver]
}' >> ./alloy-config.alloy
```{{exec}}

Reload Alloy with this config change:

```bash
curl -X POST http://localhost:12345/-/reload
```{{exec}}

After reloading Alloy, we can see the new component in the Alloy UI:
[http://localhost:12345]({{TRAFFIC_HOST1_12345}})

## Generating logs

Make sure to generate some logs using the Carnivorous Greenhouse application. You can do this by going to the Carnivorous Greenhouse UI:
[http://localhost:5005]({{TRAFFIC_HOST1_5005}})

## Viewing in Grafana

Now that we have the logs being scraped and sent to Loki, we can view them in Grafana. We can do this by going to the Explore section in Grafana and querying the logs:
[http://localhost:3000/explore]({{TRAFFIC_HOST1_3000}}/explore)



