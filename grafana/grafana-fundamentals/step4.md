# Add a logging data source

Grafana supports log data sources, like [Loki](https://grafana.com/oss/loki/). Just like for metrics, you first need to add your data source to Grafana.

1. Click the menu icon and, in the sidebar, click **Connections** and then **Data sources**.

1. Click **+ Add new data source**.

1. In the list of data sources, click **Loki**.

1. In the URL box, enter [http://loki:3100](http://loki:3100).

1. Scroll to the bottom of the page and click **Save & Test** to save your changes.

You should see the message “Data source successfully connected.” Loki is now available as a data source in Grafana.
