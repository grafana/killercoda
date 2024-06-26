# Configure {{% param “PRODUCT_NAME” %}}

Once the local Grafana instance is set up, the next step is to configure {{< param “PRODUCT_NAME” >}}.
You use components in the `config.alloy`{{copy}} file to tell {{< param “PRODUCT_NAME” >}} which logs you want to scrape, how you want to process that data, and where you want the data sent.

The examples run on a single host so that you can run them on your laptop or in a Virtual Machine.
You can try the examples using a `config.alloy`{{copy}} file and experiment with the examples yourself.

For the following steps, create a file called `config.alloy`{{copy}} in your current working directory.
If you have enabled the {{< param “PRODUCT_NAME” >}} UI, you can “hot reload” a configuration from a file.
In a later step, you will copy this file to where {{< param “PRODUCT_NAME” >}} will pick it up, and be able to reload without restarting the system service.

## First component: Log files

Paste this component into the top of the `config.alloy`{{copy}} file:

```alloy
local.file_match "local_files" {
    path_targets = [{"__path__" = "/var/log/*.log"}]
    sync_period = "5s"
}
```{{copy}}

This component creates a [local.file_match](https://grafana.com/../../reference/components/local.file_match/) component named `local_files`{{copy}} with an attribute that tells {{< param “PRODUCT_NAME” >}} which files to source, and to check for new files every 5 seconds.

## Second component: Scraping

Paste this component next in the `config.alloy`{{copy}} file:

```alloy
loki.source.file "log_scrape" {
   targets    = local.file_match.local_files.targets
   forward_to = [loki.process.filter_logs.receiver]
   tail_from_end = true
}
```{{copy}}

This configuration creates a [loki.source.file](https://grafana.com/../../reference/components/loki.source.file/) component named `log_scrape`{{copy}}, and shows the pipeline concept of {{< param “PRODUCT_NAME” >}} in action. The `log_scrape`{{copy}} component does the following:

1. It connects to the `local_files`{{copy}} component (its “source” or target).

1. It forwards the logs it scrapes to the “receiver” of another component called `filter_logs`{{copy}} that you will define next.

1. It provides extra attributes and options, in this case, you will tail log files from the end and not ingest the entire past history.

## Third component: Filter non-essential logs

Filtering non-essential logs before sending them to a data source can help you manage log volumes to reduce costs. The filtering strategy of each organization will differ as they have different monitoring needs and setups.

The following example demonstrates filtering out or dropping logs before sending them to Loki.

Paste this component next in the `config.alloy`{{copy}} file:

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

1. `loki.process`{{copy}} is a component that allows you to transform, filter, parse, and enrich log data.
   Within this component, you can define one or more processing stages to specify how you would like to process log entries before they are stored or forwarded.

1. In this example, you create a `loki.process`{{copy}} component named “filter_logs”.
   This component receives scraped log entries from the `log_scrape`{{copy}} component you created in the previous step.

1. There are many ways to [transform, filter, parse, and enrich log data](https://grafana.com/../../reference/components/loki.process/). In this example, you use the `stage.drop`{{copy}} block to drop log entries based on specified criteria.

1. You set the `source`{{copy}} parameter equal to an empty string to denote that scraped logs from the default source, the `log_scrape`{{copy}} component, will be processed.

1. You set the `expression`{{copy}} parameter equal to the log message that is not relevant to the use case.
   The log message “.*Connection closed by authenticating user root” was chosen to demonstrate how to use the `stage.drop`{{copy}} block.

1. You can include an optional string label  `drop_counter_reason`{{copy}} to show the rationale for dropping log entries.
   You can use this label to categorize and count the drops to track and analyze the reasons for dropping logs.

1. You use the `forward_to`{{copy}} parameter to specify where to send the processed logs.
   In this case, you will send the processed logs to a component you will create next called `grafana_loki`{{copy}}.

Check out the following [tutorial](https://grafana.com/./processing-logs/) and the [`loki.process`{{copy}} documentation](https://grafana.com/../../reference/components/loki.process/) for more comprehensive information on processing logs.
