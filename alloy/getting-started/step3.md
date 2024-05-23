# Step 2: Scraping System Logs

We are going to start by building out the Grafana Alloy config. To start we going to collect metrics from our local machine. 

Lets create a new `config.alloy` file and add the following:


1. Add the following to the `config.alloy` file. To do this open Vscode and select the `config.alloy` file (this needs to be explained to the user):
```json
loki.write "grafana_loki" {
  endpoint {
    url = "http://localhost:3100/loki/api/v1/push"

    basic_auth {
      username = "admin"
      password = "admin"
    }
  }
}

local.file_match "local_files" {
    path_targets = [{"__path__" = "/var/log/*"}]
    sync_period = "5s"

}

loki.source.file "log_scrape" {
    targets    = local.file_match.local_files.targets
    forward_to = [loki.write.grafana_loki.receiver]
    tail_from_end = true
}
```{{copy}}

3. Save the file.


4. Reload Alloy with this config change:

    ```bash
    curl -X POST http://localhost:12345/-/reload
    ```{{exec}}

5. After reloading Alloy, we can see the new component in the Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})

6. Finaly lets check Grafana to see if the logs are being scraped.
   [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
