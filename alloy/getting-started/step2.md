# Step 2: Building out the Grafana Alloy config

We are going to start by building out the Grafana Alloy config. To start we going to collect metrics from our local machine. 

Lets create a new `config.alloy` file and add the following:

1. Create a new `config.alloy` file in the root of the project.
   ```bash
    touch config.alloy
    ```

2. Add the following to the `config.alloy` file:
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
   ```

5. Reload Alloy with this config change:

    ```bash
    curl -X POST http://localhost:12345/-/reload
    ```{{exec}}

5. After reloading Alloy, we can see the new component in the Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})


