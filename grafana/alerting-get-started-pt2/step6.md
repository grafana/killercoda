# Create an alert rule

1. Navigate to **Alerting > Alert rules**.

1. Click **New alert rule**.

# Enter an alert rule name

Make it short and descriptive as this will appear in your alert notification. For instance, `web-traffic`{{copy}}.

# Define query and alert condition

In this section, we use the default options for Grafana-managed alert rule creation. The default options let us define the query, a expression (used to manipulate the data – the `WHEN`{{copy}} field in the UI), and the condition that must be met for the alert to be triggered (in default mode is the threshold).

1. Select **TestData** data source from the drop-down menu.

1. From **Scenario** select **CSV Content**.

1. In the Query editor, switch to **Code** mode by clicking the button on the right.

1. Copy in the following CSV data:

   ```
   device,views
   desktop,1200
   mobile,900
   ```{{copy}}

   The above CSV data simulates a data source returning multiple time series, each leading to the creation of an alert instance for that specific time series. Note that the data returned matches the example in the [Alert instance](https://grafana.com#alert-instances) section.

1. In the **Alert condition** section:

   - Keep `Last`{{copy}} as the value for the reducer function (`WHEN`{{copy}}), and `1000`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview** to run the queries.

It should return two series.`desktop`{{copy}} in Firing state, and `mobile`{{copy}} in Normal state. The values `1`{{copy}}, and `0`{{copy}} mean that the condition is either `true`{{copy}} or `false`{{copy}}.

![Screenshot showing a preview of a query in Grafana that returns two alert instances, including the query results and relevant alert details](https://grafana.com/media/docs/alerting/firing-instances.png)
