# Create an alert rule

Next, we'll establish an [alert rule](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rule-evaluation/) within Grafana Alerting to notify us whenever alert rules are triggered and resolved.

1. In Grafana, **navigate to Alerting** > **Alert rules**.
1. Click on **New alert rule**.
1. Enter alert rule name for your alert rule. Make it short and descriptive as this appears in your alert notification. For instance, **web-requests-logs**

## Define query and alert condition

In this section, we use the default options for Grafana-managed alert rule creation. The default options let us define the query, an expression (used to manipulate the data -- the `WHEN`{{copy}} field in the UI), and the condition that must be met for the alert to be triggered (in default mode is the threshold).

1. Select the **Loki** data source from the drop-down.
1. The query editor opens in **Code** mode by default. If it's showing **Builder** mode instead, click **Code** on the right.
1. Paste the query below.

   ```
   sum by (message)(count_over_time({filename="/var/log/web_requests.log"} != "status=200" | pattern "<_> <message> duration<_>" [10m]))
   ```{{copy}}

   This query counts the number of log lines with a status code that is not 200 (OK), then sum the result set by message type using an **instant query** and the time interval indicated in brackets. It uses the LogQL pattern parser to add a new label called `message`{{copy}} that contains the level, method, url, and status from the log line.

   You can use the **explain query** toggle button for a full explanation of the query syntax. The optional log-generating script creates a sample log line similar to the one below:

   ```
   2023-04-22T02:49:32.562825+00:00 level=info method=GET url=test.com status=200 duration=171ms
   ```{{copy}}

   If you're using your own logs, modify the LogQL query to match your own log message. Refer to the Loki docs to understand the [pattern parser](https://grafana.com/docs/loki/latest/logql/log_queries/#pattern).
1. In the **Alert condition** section, confirm the condition reads **WHEN QUERY IS ABOVE** `0`{{copy}}. This is the default, and `0`{{copy}} is the threshold above which the alert rule should trigger.
1. Click **Preview alert rule condition** to run the query.

   It should return alert instances from log lines with a status code that is not 200 (OK), and that has met the alert condition. The condition for the alert rule to fire is any occurrence that goes over the threshold of `0`{{copy}}. Since the Loki query has returned more than zero alert instances, the alert rule is `Firing`{{copy}}.

   ![Preview of a firing alert instances](https://grafana.com/media/docs/alerting/firing-loki-alert-rule.png)

## Set evaluation behavior

The [alert rule evaluation](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/rule-evaluation/) defines the conditions under which an alert rule triggers, based on the following settings:

- **Evaluation group**: every alert rule is assigned to an evaluation group. You can assign the alert rule to an existing evaluation group or create a new one.
- **Evaluation interval**: determines how frequently the alert rule is checked. For instance, the evaluation may occur every 10s, 30s, 1m, 10m, etc.
- **Pending period**: how long the condition must be met to trigger the alert rule.
- **Keep firing for**: defines how long an alert should remain in the Firing state after the alert condition stops being true. During this time, the alert enters a _Recovering_ state, suppressing additional notifications but keeping the alert active. It helps prevent alert flapping, where alerts rapidly switch between firing and resolved due to noisy or unstable metrics.

To set up the evaluation:

1. In **Folder**, click **+ New folder** and enter a name. For example: _web-server-alerts_. This folder contains our alerts.
1. In **Evaluation group and interval**, click **+ New evaluation group**.
1. In **Evaluation group name**, enter _1m-evaluation_.
1. Keep **Evaluation interval** set to `1m`{{copy}} (1 minute), so the alert rule is evaluated every minute, and then click **Create**.
1. Set the pending period to, `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met.
1. Set **Keep firing for** to, `0s`{{copy}}, so the alert stops firing immediately after the condition is no longer true. Use this when you want alerts to be resolved as soon as the system is healthy again.

## Configure labels and notifications

Choose the contact point where you want to receive your alert notifications.

1. Under **Contact point**, select **Webhook** from the drop-down menu.
1. Click **Save** at the bottom of the page.
