# Create an alert rule

1. Navigate to **Alerting > Alert rules**.

1. Click **New alert rule**.

# Enter an alert rule name

Make it short and descriptive as this will appear in your alert notification. For instance, `web-traffic`{{copy}}.

# Define query and alert condition

In this section, we use the **Advanced options** for Grafana-managed alert rule creation. The advanced options let us define queries, expressions (used to manipulate the data), and the condition that must be met for the alert to be triggered.

1. Toggle **Advanced options** to view additional configuration options.

1. Select **TestData** data source from the drop-down menu.

1. From **Scenario** select **CSV Content**.

1. Copy in the following CSV data:

```
device,views
desktop,1200
mobile,900
```{{copy}}

The above CSV data simulates a data source returning multiple time series, each leading to the creation of an alert instance for that specific time series. Note that the data returned matches the example in the [Alert instance](https://grafana.com#alert-instances) section.

1. Remove the ‘B’ **Reduce expression** (click the bin icon). The Reduce expression is default, and in this case, is not required since the queried data is already reduced. Note that the Threshold expression is now your **Alert condition**.

1. In the ‘C’ **Threshold expression**:
   - Change the **Input** to ‘**A**’ to select the data source.

   - Enter `1000`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview** to run the queries.

It should return two series.`desktop`{{copy}} in Firing state, and `mobile`{{copy}} in Normal state. The values `1`{{copy}}, and `0`{{copy}} mean that the condition is either `true`{{copy}} or `false`{{copy}}.

![Screenshot showing a preview of a query in Grafana that returns two alert instances, including the query results and relevant alert details](https://grafana.com/media/docs/alerting/get-started-expression-instances.png)
