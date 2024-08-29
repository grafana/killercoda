# Create an alert ruke

Next, we’ll establish an [alert rule](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/) within Grafana Alerting to notify us whenever alert rules are triggered and resolved.

1. In Grafana, **navigate to Alerting** > **Alert rules**.

1. Click on **New alert rule**.

1. Enter alert rule name for your alert rule. Make it short and descriptive as this will appear in your alert notification. For instance, **web-requests-logs**

## Define query and alert condition

In this section, we define queries, expressions (used to manipulate the data), and the condition that must be met for the alert to be triggered.

1. Select the **Loki** datasource from the drop-down.

1. In the Query editor, switch to Code mode by clicking the button on the right.

1. Paste the query below.

   ```
   sum by (message)(count_over_time({filename="/var/log/web_requests.log"} != "status=200" | pattern "<_> <message> duration<_>" [10m]))
   ```{{copy}}

This query will count the number of log lines with a status code that is not 200 (OK), then sum the result set by message type using an **instant query** and the time interval indicated in brackets. It uses the LogQL pattern parser to add a new label called `message`{{copy}} that contains the level, method, url, and status from the log line.

You can use the **explain query** toggle button for a full explanation of the query syntax. The optional log-generating script creates a sample log line similar to the one below:

```
2023-04-22T02:49:32.562825+00:00 level=info method=GET url=test.com status=200 duration=171ms
```{{copy}}

If you’re using your own logs, modify the LogQL query to match your own log message. Refer to the Loki docs to understand the [pattern parser](https://grafana.com/docs/loki/latest/logql/log_queries/#pattern).

1. Remove the ‘B’ **Reduce expression** (click the bin icon). The Reduce expression comes by default, and in this case, it is not needed since the queried data is already reduced. Note that the Threshold expression is now your **Alert condition**.

1. In the ‘C’ **Threshold expression**:

   - Change the **Input** to **‘A’** to select the data source.

   - Enter `0`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview** to run the queries.

   It should return a single sample with the value 1 at the current timestamp. And, since `1`{{copy}} is above `0`{{copy}}, the alert condition has been met, and the alert rule state is `Firing`{{copy}}.

   ![Preview of a firing alert instances](https://grafana.com/media/docs/alerting/expression-loki-alert.png)

## Set evaluation behavior

An [evaluation group](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/rule-evaluation/) defines when an alert rule fires, and it’s based on two settings:

- **Evaluation group**: how frequently the alert rule is evaluated.

- **Evaluation interval**: how long the condition must be met to start firing. This allows your data time to stabilize before triggering an alert, helping to reduce the frequency of unnecessary notifications.

To set up the evaluation:

1. In **Folder**, click **+ New folder** and enter a name. For example: _loki-alerts_. This folder will contain our alerts.

1. In the **Evaluation group**, repeat the above step to create a new evaluation group. We will name it _1m-evaluation_.

1. Choose an **Evaluation interval** (how often the alert will be evaluated).
   For example, every `1m`{{copy}} (1 minute).

1. Set the pending period to, `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met.

## Configure labels and notifications

Choose the contact point where you want to receive your alert notifications.

1. Under **Contact point**, select **Webhook** from the drop-down menu.

1. Click **Save rule and exit** at the top right corner.
