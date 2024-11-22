# Create an alert

Next, we establish an [alert rule](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/create-grafana-managed-rule/) within Grafana Alerting to notify us whenever alert rules are triggered and resolved.

1. In Grafana, **navigate to Alerting** > **Alert rules**. Click on **New alert rule**.

1. Enter alert rule name for your alert rule. Make it short and descriptive as this appears in your alert notification. For instance, **database-metrics**

## Define query and alert condition

In this section, we use the default options for Grafana-managed alert rule creation. The default options let us define the query, a expression (used to manipulate the data – the `WHEN`{{copy}} field in the UI), and the condition that must be met for the alert to be triggered (in default mode is the threshold).

Grafana includes a [test data source](https://grafana.com/docs/grafana/latest/datasources/testdata/) that creates simulated time series data. This data source is included in the demo environment for this tutorial. If you’re working in Grafana Cloud or your own local Grafana instance, you can add the data source through the **Connections** menu.

1. Select the **TestData** data source from the drop-down menu.

1. In the **Alert condition** section:

   - Keep `Last`{{copy}} as the value for the reducer function (`WHEN`{{copy}}), and `0`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview alert rule condition** to run the query.

   It should return random time series data. The alert rule state should be `Firing`{{copy}}.

   ![A preview of a firing alert](https://grafana.com/media/docs/alerting/random-walk-firing-alert-rule.png)

## Set evaluation behavior

The [alert rule evaluation](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/rule-evaluation/) defines the conditions under which an alert rule triggers, based on the following settings:

- **Evaluation group**: every alert rule is assigned to an evaluation group. You can assign the alert rule to an existing evaluation group or create a new one.

- **Evaluation interval**: determines how frequently the alert rule is checked. For instance, the evaluation may occur every 10s, 30s, 1m, 10m, etc.

- **Pending period**: how long the condition must be met to trigger the alert rule.

To set up the evaluation:

1. In **Folder**, click **+ New folder** and enter a name. For example: _metric-alerts_. This folder contains our alerts.

1. In the **Evaluation group**, repeat the above step to create a new evaluation group. Name it _1m-evaluation_.

1. Choose an **Evaluation interval** (how often the alert are evaluated).
   For example, every `1m`{{copy}} (1 minute).

1. Set the pending period to, `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met.

## Configure labels and notifications

Choose the contact point where you want to receive your alert notifications.

1. Under **Contact point**, select **Webhook** from the drop-down menu.

1. Click **Save rule and exit** at the top right corner.
