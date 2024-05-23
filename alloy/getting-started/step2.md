# Step 2: Scraping system metrics

We are going to start by building out the Grafana Alloy config. To start we going to collect metrics from our local machine. 

Lets create a new `config.alloy` file and add the following:

1. Create a new `config.alloy` file in the root of the project.
   ```bash
    touch config.alloy
    ```{{exec}}

2. Add the following to the `config.alloy` file. To do this open Vscode and select the `config.alloy` file (this needs to be explained to the user):
```json
prometheus.exporter.unix "local_system" { }

// Configure a prometheus.scrape component to collect unix metrics.
prometheus.scrape "scrape_metrics" {
  targets    = prometheus.exporter.unix.local_system.targets
  forward_to = [prometheus.remote_write.metrics_service.receiver]
  scrape_interval = "10s"
}

prometheus.remote_write "metrics_service" {
    endpoint {
        url = "http://localhost:9090/api/v1/write"

        basic_auth {
            username = "admin"
            password = "admin"
        }
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


6. After reloading Alloy, we can see the new component in the Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})

7. Finaly lets check Grafana to see if the metrics are being scraped.
   [http://localhost:3000]({{TRAFFIC_HOST1_3000}})
