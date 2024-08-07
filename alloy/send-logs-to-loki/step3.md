# Configure Alloy

After the local Grafana instance is set up, the next step is to configure Alloy.
You use components in the `config.alloy`{{copy}} file to tell Alloy which logs you want to scrape, how you want to process that data, and where you want the data sent.

The examples run on a single host so that you can run them on your laptop or in a Virtual Machine.
You can try the examples using a `config.alloy`{{copy}} file and experiment with the examples.

## First component: Log files

`local.file_match`{{copy}} is a component that tells Alloy which log files to source:

1. Create a file called `config.alloy`{{copy}} in your current working directory:

   ```bash
    touch config.alloy
   ```{{exec}}

1. Copy and paste the following component configuration at the top of the file:

   ```alloy
   ```{{copy}}

local.file_match “local_files” {
path_targets = [{"**path**" = “/var/log/*.log”}]
sync_period = “5s”
}

```

This configuration creates a [local.file_match][] component named `local_files` which does the following:

* It tells Alloy which files to source.
* It checks for new files every 5 seconds.

### Second component: Scraping

The next component scrapes the logs from the log files you specified in the first component:

1. Copy and paste the following component configuration below the previous component in your `config.alloy` file:
```alloy
loki.source.file "log_scrape" {
  targets    = local.file_match.local_files.targets
  forward_to = [loki.process.filter_logs.receiver]
  tail_from_end = true
}
```{{copy}}

This configuration creates a [loki.source.file](https://grafana.com/../../reference/components/loki/loki.source.file/) component named `log_scrape`{{copy}} which does the following:

- It connects to the `local_files`{{copy}} component as its source or target.

- It forwards the logs it scrapes to the receiver of another component called `filter_logs`{{copy}}.

- It provides extra attributes and options to tail the log files from the end so you don’t ingest the entire log file history.

## Third component: Filter non-essential logs

Filtering non-essential logs before sending them to a data source can help you manage log volumes to reduce costs.

The following example demonstrates how you can filter out or drop logs before sending them to Loki.

1. Copy and paste the following component configuration below the previous component in your `config.alloy`{{copy}} file:

```alloy
loki.process "filter_logs" {
  stage.drop {
      source = ""
      expression  = ".*Connection closed by authenticating user root"
      drop_counter_reason = "noisy"
    }
  forward_to = [loki.write.grafana_loki.receiver]
  }
```{{copy}}

The `loki.process`{{copy}} component allows you to transform, filter, parse, and enrich log data.
Within this component, you can define one or more processing stages to specify how you would like to process log entries before they’re stored or forwarded.

This configuration creates a [loki.process](https://grafana.com/../../reference/components/loki/loki.process/) component named `filter_logs`{{copy}} which does the following:

- It receives scraped log entries from the default `log_scrape`{{copy}} component.

- It uses the `stage.drop`{{copy}} block to define what to drop from the scraped logs.

- It uses the `expression`{{copy}} parameter to identify the specific log entries to drop.

- It uses an optional string label `drop_counter_reason`{{copy}} to show the reason for dropping the log entries.

- It forwards the processed logs to the receiver of another component called `grafana_loki`{{copy}}.

The [`loki.process`{{copy}} documentation](https://grafana.com/../../reference/components/loki/loki.process/) provides more comprehensive information on processing logs.

## Fourth component: Write logs to Loki

Lastly, you need to configure a component to write the processed logs to Loki:

1. Copy and paste this component configuration below the previous component in your `config.alloy`{{copy}} file:

```alloy
loki.write "grafana_loki" {
  endpoint {
    url = "http://localhost:3100/loki/api/v1/push"

    // basic_auth {
    //  username = "admin"
    //  password = "admin"
    // }
  }
}
```{{copy}}

This final component creates a [`loki.write`{{copy}}][] component named `grafana_loki`{{copy}} that points to `http://localhost:3100/loki/api/v1/push`{{copy}}.

This completes the simple configuration pipeline.

> The `basic_auth` block is commented out because the local `docker-compose` stack doesn't require it. It's included in this example to show how you can configure authorization for other environments. For further authorization options, refer to the [loki.write][] component reference.
With this configuration, Alloy connects directly to the Loki instance running in the Docker container.
