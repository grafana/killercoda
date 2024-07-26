# Trigger and resolve an alert

Now that the alert rule has been configured, you should receive alert [notifications](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/state-and-health/#notifications) in the contact point whenever alerts trigger and get resolved.

## Trigger an alert

Since the alert rule that you have created has been configured to always fire, once the evaluation interval has concluded, you should receive an alert notification in the Webhook endpoint.

![Firing alert notification details](https://grafana.com/media/docs/alerting/alerting-webhook-firing-alert.png)

The alert notification details show that the alert rule state is Firing , and it includes the value that made the rule trigger by exceeding the threshold of the alert rule condition. The notification also includes links to see the alert rule details, and another link to add a [Silence](http://grafana.com/docs/grafana/next/alerting/configure-notifications/create-silence/) to it.

## Resolve an alert

To see how a resolved alert notification looks like, you can modify the current alert rule threshold.

To edit the Alert rule:

1. **Navigate to Alerting** > **Alert rules**.

1. Click on the metric-alerts folder to display the alert that you created earlier

1. Click the **edit** button on the right hand side of the screen

1. Increment the Threshold expression to 1.

1. Click **Save rule and exit**.

By incrementing the threshold, the condition is no longer met, and after the evaluation interval has concluded (1 minute approx.), you should receive an alert notification with status **“Resolved”**.
