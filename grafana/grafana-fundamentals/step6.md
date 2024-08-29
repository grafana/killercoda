# Build a dashboard

A _dashboard_ gives you an at-a-glance view of your data and lets you track metrics through different visualizations.

Dashboards consist of _panels_, each representing a part of the story you want your dashboard to tell.

Every panel consists of a _query_ and a _visualization_. The query defines _what_ data you want to display, whereas the visualization defines _how_ the data is displayed.

1. Click the menu icon and, in the sidebar, click **Dashboards**.

1. On the **Dashboards** page, click **New** in top right corner and select **New Dashboard** in the drop-down.

1. Click **+ Add visualization**.

1. In the modal that opens, select the Prometheus data source that you just added.

1. In the **Query** tab below the graph, enter the query from earlier and then press Shift + Enter:

   ```
   sum(rate(tns_request_duration_seconds_count[5m])) by(route)
   ```{{copy}}

1. In the panel editor on the right, under **Panel options**, change the panel title to “Traffic”.

1. Click **Apply** in the top-right corner to save the panel and go back to the dashboard view.

1. Click the **Save dashboard** (disk) icon at the top of the dashboard to save your dashboard.

1. Enter a name in the **Dashboard name** field and then click **Save**.

   You should now have a panel added to your dashboard.

   ![A panel in a Grafana dashboard](https://grafana.com/media/tutorials/grafana-fundamentals-dashboard.png)
