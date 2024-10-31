# Create an alert

Next, we establish an [alert rule](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/create-grafana-managed-rule/) within Grafana Alerting to notify us whenever alert rules are triggered and resolved.

1. In Grafana, **navigate to Alerting** > **Alert rules**. Click on **New alert rule**.

1. Enter alert rule name for your alert rule. Make it short and descriptive as this will appear in your alert notification. For instance, **database-metrics**

## Define query and alert condition

In this section, we use the default options for Grafana-managed alert rule creation. The default options let us define the query, a expression (used to manipulate the data – the `WHEN`{{copy}} field in the UI), and the condition that must be met for the alert to be triggered (in default mode is the threshold).

1. Select the **Prometheus** data source from the drop-down menu.

1. In the Query editor, switch to **Code** mode by clicking the button on the right.

1. Enter the following query:

   ```promql
   vector(1)
   ```{{copy}}

   In Prometheus, `vector(1)`{{copy}} is a special type of PromQL query that generates a constant vector. This is useful in testing and query manipulation, where you might need a constant value for calculations or comparisons. This query will allow you to create an alert rule that will be always firing.

1. In the **Alert condition** section:

   - Keep `Last`{{copy}} as the value for the reducer function (`WHEN`{{copy}}), and `0`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview alert rule condition** to run the query.

   It should return a single sample with the value 1 at the current timestamp. And, since `1`{{copy}} is above `0`{{copy}}, the alert condition has been met, and the alert rule state is `Firing`{{copy}}.

   ![A preview of a firing alert](https://grafana.com/media/docs/alerting/firing-alert-preview.png)

## Set evaluation behavior

The [alert rule evaluation](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/rule-evaluation/) defines the conditions under which an alert rule triggers, based on the following settings:

- **Evaluation group**: every alert rule is assigned to an evaluation group. You can assign the alert rule to an existing evaluation group or create a new one.

- **Evaluation interval**: determines how frequently the alert rule is checked. For instance, the evaluation may occur every 10s, 30s, 1m, 10m, etc.

- **Pending period**: how long the condition must be met to trigger the alert rule.

To set up the evaluation:

1. In **Folder**, click **+ New folder** and enter a name. For example: _metric-alerts_. This folder will contain our alerts.

1. In the **Evaluation group**, repeat the above step to create a new evaluation group. We will name it _1m-evaluation_.

1. Choose an **Evaluation interval** (how often the alert will be evaluated).
   For example, every `1m`{{copy}} (1 minute).

1. Set the pending period to, `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met.

## Configure labels and notifications

Choose the contact point where you want to receive your alert notifications.

1. Under **Contact point**, select **Webhook** from the drop-down menu.

1. Click **Save rule and exit** at the top right corner.
