
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


