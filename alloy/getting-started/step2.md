# Step 2: Scraping system metrics

We are going to start by building out the Grafana Alloy config. To start we going to collect metrics from our local machine. 

Lets create a new `config.alloy` file and add the following:

1. Create a new `config.alloy` file in the root of the project.
   ```bash
    touch config.alloy
    ```

2. Add the following to the `config.alloy` file (this needs to be explained to the user):
```json
prometheus.exporter.self "integrations_alloy" { }

discovery.relabel "integrations_alloy" {
  targets = prometheus.exporter.self.integrations_alloy.targets

  rule {
    target_label = "instance"
    replacement  = constants.hostname
  }

  rule {
    target_label = "alloy_hostname"
    replacement  = constants.hostname
  }

  rule {
    target_label = "job"
    replacement  = "integrations/alloy-check"
  }
}
```{{copy}}

3. Save the file.

4. Lets copy the `config.alloy` file to the Alloy config directory.
   ```bash
   sudo cp config.alloy /etc/alloy/config.alloy
   ```{{exec}}

5. Reload Alloy with this config change:

    ```bash
    curl -X POST http://localhost:12345/-/reload
    ```{{exec}}

5. After reloading Alloy, we can see the new component in the Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})


6. Lets scrape some metrics ....
```json
prometheus.scrape "integrations_alloy" {
  targets    = discovery.relabel.integrations_alloy.output
  forward_to = [prometheus.relabel.integrations_alloy.receiver]  

  scrape_interval = "5s"
}

prometheus.relabel "integrations_alloy" {
  forward_to = [prometheus.remote_write.metrics_service.receiver]

  rule {
    source_labels = ["__name__"]
    regex         = "(prometheus_target_sync_length_seconds_sum|prometheus_target_scrapes_.*|prometheus_target_interval.*|prometheus_sd_discovered_targets|alloy_build.*|prometheus_remote_write_wal_samples_appended_total|process_start_time_seconds)"
    action        = "keep"
  }
}
```{{copy}}


7. Now lets send these metrics to Prometheus:
```json
prometheus.remote_write "metrics_service" {
  endpoint {
    url = "http://localhost:9090/api/v1/pushs"
  }
}
```{{copy}}

8. Lets copy the `config.alloy` file to the Alloy config directory.
   ```bash
   sudo cp config.alloy /etc/alloy/config.alloy
   ```{{exec}}

9. Reload Alloy with this config change:

    ```bash
    curl -X POST http://localhost:12345/-/reload
    ```{{exec}}

10. After reloading Alloy, we can see the new component in the Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})
