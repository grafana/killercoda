# Annotate events

When things go bad, it often helps if you understand the context in which the failure occurred. Time of last deploy, system changes, or database migration can offer insight into what might have caused an outage. Annotations allow you to represent such events directly on your graphs.

In the next part of the tutorial, we will simulate some common use cases that someone would add annotations for.

1. To manually add an annotation, click anywhere in your graph, then click **Add annotation**.
   Note: you might need to save the dashboard first.

1. In **Description**, enter **Migrated user database**.

1. Click **Save**.

   Grafana adds your annotation to the graph. Hover your mouse over the base of the annotation to read the text.

Grafana also lets you annotate a time interval, with _region annotations_.

Add a region annotation:

1. Press Ctrl (or Cmd on macOS) and hold, then click and drag across the graph to select an area.

1. In **Description**, enter **Performed load tests**.

1. In **Tags**, enter **testing**.

1. Click **Save**.

## Using annotations to correlate logs with metrics

Manually annotating your dashboard is fine for those single events. For regularly occurring events, such as deploying a new release, Grafana supports querying annotations from one of your data sources. Let’s create an annotation using the Loki data source we added earlier.

1. At the top of the dashboard, click the **Dashboard settings** (gear) icon.

1. Go to **Annotations** and click **Add annotation query**.

1. In **Name**, enter **Errors**.

1. In **Data source**, select **Loki**.

1. In **Query**, enter the following query:

   ```
   {filename="/var/log/tns-app.log"} |= "error"
   ```{{copy}}

1. Click **Apply**. Grafana displays the Annotations list, with your new annotation.

1. Click on your dashboard name to return to your dashboard.

1. At the top of your dashboard, there is now a toggle to display the results of the newly created annotation query. Press it if it’s not already enabled.

1. Click the **Save dashboard** icon to save the changes.

1. To test the changes, go back to the [sample application]({{TRAFFIC_HOST1_8081}}), post a new link without a URL to generate an error in your browser that says `empty url`{{copy}}.

The log lines returned by your query are now displayed as annotations in the graph.

![A panel in a Grafana dashboard with log queries from Loki displayed as annotations](https://grafana.com/media/tutorials/annotations-grafana-dashboard.png)

Being able to combine data from multiple data sources in one graph allows you to correlate information from both Prometheus and Loki.

Annotations also work very well alongside alert rules. In the next and final section, we will set up an alert rules for our app `grafana.news`{{copy}} and then we will trigger it. This will provide a quick intro to our new Alerting platform.
