# Configure Alloy

In this tutorial, you configure Alloy to collect metrics and send them to Prometheus.

You add components to your `config.alloy`{{copy}} file to tell Alloy which metrics you want to scrape, how you want to process that data, and where you want the data sent.

The following steps build on the `config.alloy`{{copy}} file you created in the previous tutorial.

> The interactive sandbox has a VSCode-like editor that allows you to access files and folders. To access this feature, click on the `Editor`{{copy}} tab. The editor also has a terminal that you can use to run commands. Since some commands assume you are within a specific directory, we recommend running the commands in `tab1`{{copy}}.

## First component: Scraping

Paste the following component configuration at the top of your `config.alloy`{{copy}} file:

```alloy
prometheus.exporter.unix "local_system" { }

prometheus.scrape "scrape_metrics" {
  targets         = prometheus.exporter.unix.local_system.targets
  forward_to      = [prometheus.relabel.filter_metrics.receiver]
  scrape_interval = "10s"
}
```{{copy}}

This configuration creates a [`prometheus.scrape`{{copy}}](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.scrape/) component named `scrape_metrics`{{copy}} which does the following:

- It connects to the `local_system`{{copy}} component as its source or target.

- It forwards the metrics it scrapes to the receiver of another component called `filter_metrics`{{copy}}.

- It tells Alloy to scrape metrics every 10 seconds.

## Second component: Filter metrics

Filtering non-essential metrics before sending them to a data source can help you reduce costs and allow you to focus on the data that matters most.

The following example demonstrates how you can filter out or drop metrics before sending them to Prometheus.

Paste the following component configuration below the previous component in your `config.alloy`{{copy}} file:

```alloy
prometheus.relabel "filter_metrics" {
  rule {
    action        = "drop"
    source_labels = ["env"]
    regex         = "dev"
  }

  forward_to = [prometheus.remote_write.metrics_service.receiver]
}
```{{copy}}

The [`prometheus.relabel`{{copy}}](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.relabel/) component is commonly used to filter Prometheus metrics or standardize the label set passed to one or more downstream receivers.
You can use this component to rewrite the label set of each metric sent to the receiver.
Within this component, you can define rule blocks to specify how you would like to process metrics before they’re stored or forwarded.

This configuration creates a [`prometheus.relabel`{{copy}}](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.relabel/) component named `filter_metrics`{{copy}} which does the following:

- It receives scraped metrics from the `scrape_metrics`{{copy}} component.

- It tells Alloy to drop metrics that have an `"env"`{{copy}} label equal to `"dev"`{{copy}}.

- It forwards the processed metrics to the receiver of another component called `metrics_service`{{copy}}.

## Third component: Write metrics to Prometheus

Paste the following component configuration below the previous component in your `config.alloy`{{copy}} file:

```alloy
prometheus.remote_write "metrics_service" {
    endpoint {
        url = "http://localhost:9090/api/v1/write"

        // basic_auth {
        //   username = "admin"
        //   password = "admin"
        // }
    }
}
```{{copy}}

This final component creates a [`prometheus.remote_write`{{copy}}](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.remote_write/) component named `metrics_service`{{copy}} that points to `http://localhost:9090/api/v1/write`{{copy}}.

This completes the simple configuration pipeline.

> The `basic_auth`{{copy}} is commented out because the local `docker-compose`{{copy}} stack doesn’t require it. It’s included in this example to show how you can configure authorization for other environments. For further authorization options, refer to the [`prometheus.remote_write`{{copy}}](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.remote_write/) component documentation.

This connects directly to the Prometheus instance running in the Docker container.
