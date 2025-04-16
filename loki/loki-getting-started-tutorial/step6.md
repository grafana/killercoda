# A look under the hood

At this point you will have a running Loki stack and a sample application generating logs. You have also queried Loki using Grafana Logs Drilldown and Grafana Explore.
In this next section we will take a look under the hood to understand how the Loki stack has been configured to collect logs, the Loki configuration file, and how the Loki data source has been configured in Grafana.

## Grafana Alloy configuration

Grafana Alloy is collecting logs from all the Docker containers and forwarding them to Loki.
It needs a configuration file to know which logs to collect and where to forward them to. Within the `loki-fundamentals`{{copy}} directory, you will find a file called `config.alloy`{{copy}}:

```alloy
// This component is responsible for discovering new containers within the Docker environment
discovery.docker "getting_started" {
	host             = "unix:///var/run/docker.sock"
	refresh_interval = "5s"
}

// This component is responsible for relabeling the discovered containers
discovery.relabel "getting_started" {
	targets = []

	rule {
		source_labels = ["__meta_docker_container_name"]
		regex         = "/(.*)"
		target_label  = "container"
	}
}

// This component is responsible for collecting logs from the discovered containers
loki.source.docker "getting_started" {
	host             = "unix:///var/run/docker.sock"
	targets          = discovery.docker.getting_started.targets
	forward_to       = [loki.process.getting_started.receiver]
	relabel_rules    = discovery.relabel.getting_started.rules
	refresh_interval = "5s"
}

// This component is responsible for processing the logs (In this case adding static labels)
loki.process "getting_started" {
    stage.static_labels {
    values = {
      env = "production",
    }
}
    forward_to = [loki.write.getting_started.receiver]
}

// This component is responsible for writing the logs to Loki
loki.write "getting_started" {
	endpoint {
		url  = "http://loki:3100/loki/api/v1/push"
	}
}

// Enables the ability to view logs in the Alloy UI in realtime
livedebugging {
  enabled = true
}
```{{copy}}

This configuration file can be viewed visually via the Alloy UI at [http://localhost:12345/graph]({{TRAFFIC_HOST1_12345}}/graph).

![Alloy UI](https://grafana.com/media/docs/loki/getting-started-alloy-ui.png)

In this view you can see the components of the Alloy configuration file and how they are connected:

- **discovery.docker**: This component queries the metadata of the Docker environment via the Docker socket and discovers new containers, as well as providing metadata about the containers.

- **discovery.relabel**: This component converts a metadata (`__meta_docker_container_name`{{copy}}) label into a Loki label (`container`{{copy}}).

- **loki.source.docker**: This component collects logs from the discovered containers and forwards them to the next component. It requests the metadata from the `discovery.docker`{{copy}} component and applies the relabeling rules from the `discovery.relabel`{{copy}} component.

- **loki.process**: This component provides stages for log transformation and extraction. In this case it adds a static label `env=production`{{copy}} to all logs.

- **loki.write**: This component writes the logs to Loki. It forwards the logs to the Loki endpoint `http://loki:3100/loki/api/v1/push`{{copy}}.

## View Logs in realtime

Grafana Alloy provides a built-in real time log viewer. This allows you to view current log entries and how they are being transformed via specific components of the pipeline.
To view live debugging mode open a browser tab and navigate to: [http://localhost:12345/debug/loki.process.getting_started]({{TRAFFIC_HOST1_12345}}/debug/loki.process.getting_started).
