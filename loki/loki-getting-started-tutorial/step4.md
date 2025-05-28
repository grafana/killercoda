# Querying logs

At this point, you have viewed logs using the Grafana Logs Drilldown feature. In many cases this will provide you with all the information you need. However, we can also manually query Loki to ask more advanced questions about the logs. This can be done via **Grafana Explore**.

1. Open a browser and navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}) to open Grafana.
1. From the Grafana main menu, click the **Explore** icon (1) to open the Explore tab.

   To learn more about Explore, refer to the [Explore](https://grafana.com/docs/grafana/latest/explore/) documentation.

   ![Grafana Explore](https://grafana.com/media/docs/loki/grafana-query-builder-v2.png)
1. From the menu in the dashboard header, select the Loki data source (2).

   This displays the Loki query editor.

   In the query editor you use the Loki query language, [LogQL](https://grafana.com/docs/loki/latest/query/), to query your logs.
   To learn more about the query editor, refer to the [query editor documentation](https://grafana.com/docs/grafana/latest/datasources/loki/query-editor/).
1. The Loki query editor has two modes (3):

   - [Builder mode](https://grafana.com/docs/grafana/latest/datasources/loki/query-editor/#builder-mode), which provides a visual query designer.
   - [Code mode](https://grafana.com/docs/grafana/latest/datasources/loki/query-editor/#code-mode), which provides a feature-rich editor for writing LogQL queries.

   Next we’ll walk through a few queries using the code view.
1. Click **Code** (3) to work in Code mode in the query editor.

   Here are some sample queries to get you started using LogQL. After copying any of these queries into the query editor, click **Run Query** (4) to execute the query.

   1. View all the log lines which have the `container`{{copy}} label value `greenhouse-main_app-1`{{copy}}:

      ```logql
      {container="greenhouse-main_app-1"}
      ```{{copy}}

      In Loki, this is a log stream.

      Loki uses [labels](https://grafana.com/docs/loki/latest/get-started/labels/) as metadata to describe log streams.

      Loki queries always start with a label selector.
      In the previous query, the label selector is `{container="greenhouse-main_app-1"}`{{copy}}.
   1. Find all the log lines in the `{container="greenhouse-main_app-1"}`{{copy}} stream that contain the string `POST`{{copy}}:

      ```logql
      {container="greenhouse-main_app-1"} |= "POST"
      ```{{copy}}
