# Create an alert rule

In this section we configure an alert rule based on our application monitoring example.

1. Go to **Alerting > Alert rules**.

1. Click **New alert rule**.

## Enter an alert rule name

Make it short and descriptive as this will appear in your alert notification. For instance, `High CPU usage - Multi-region`{{copy}}.

## Define query and alert condition

In this section, we use the default options for Grafana-managed alert rule creation. The default options let us define the query, a expression (used to manipulate the data – the `WHEN`{{copy}} field in the UI), and the condition that must be met for the alert to be triggered (in default mode is the threshold).

Grafana includes a [test data source](https://grafana.com/docs/grafana/latest/datasources/testdata/) that creates simulated time series data. This data source is included in the demo environment for this tutorial. If you’re working in Grafana Cloud or your own local Grafana instance, you can add the data source through the **Connections** menu.

1. Select **TestData** data source from the drop-down menu.

1. From **Scenario** select **CSV Content**.

1. Copy in the following CSV data:

   - Select **TestData** as the data source.

   - Set **Scenario** to **CSV Content**.

   - Use the following CSV data:

     ```csv
     region,cpu-usage,service,instance
     us-west,35,web-server-1,server-01
     us-west,81,web-server-1,server-02
     us-east,79,web-server-2,server-03
     us-east,52,web-server-2,server-04
     us-west,45,db-server-1,server-05
     us-east,77,db-server-2,server-06
     us-west,82,db-server-1,server-07
     us-east,93,db-server-2,server-08
     ```{{copy}}

   The returned data simulates a data source returning multiple time series, each leading to the creation of an alert instance for that specific time series.

1. In the **Alert condition** section:

   - Keep `Last`{{copy}} as the value for the reducer function (`WHEN`{{copy}}), and `75`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview alert rule condition** to run the queries.

   It should return 5 series in Firing state, two firing instances from the us-west region, and three from the us-east region.

   ![Preview of a query returning alert instances.](https://grafana.com/media/docs/alerting/regions-alert-instance-preview.png)

## Set evaluation behavior

Every alert rule is assigned to an evaluation group. You can assign the alert rule to an existing evaluation group or create a new one.

1. In **Folder**, click **+ New folder** and enter a name. For example: `Multi-region CPU alerts`{{copy}}. This folder contains our alert rules.

1. In the **Evaluation group**, repeat the above step to create a new evaluation group. Name it `Multi-region CPU group`{{copy}}.

1. Choose an **Evaluation interval** (how often the alert are evaluated). Choose `1m`{{copy}}.

   The evaluation interval of 1 minute allows Grafana to detect changes quickly, while the longer **Group wait** (from our notification policy) and **Group interval** (inherited from the Default notification policy) allow for efficient grouping of alerts and minimize unnecessary notifications.

1. Set the pending period to `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met (this minimizes the waiting time for the demonstration).

## Configure labels and notifications

Choose the notification policy where you want to receive your alert notifications.

1. Select **Use notification policy**.

1. Click **Preview routing** to ensure correct matching.

   ![Preview of alert instance routing with the region label matcher](https://grafana.com/media/docs/alerting/region-notification-policy-routing-preview.png)

   The preview shows that the region label from our data source is successfully matching the notification policies that we created earlier thanks to the label matcher that we configured.

1. Click **Save rule and exit**.
