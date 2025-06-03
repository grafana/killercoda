# Add folders and labels

1. In **Folder**, click **+ New folder** and enter a name. For example: `system-metrics`{{copy}} . This folder contains our alert rules.

# Set evaluation behaviour

1. Click + **New evaluation group**. Name it `system-usage`{{copy}}.
1. Choose an **Evaluation interval** (how often the alert will be evaluated). Choose `1m`{{copy}}.
1. Set the **pending period** to `0s`{{copy}} (None), so the alert rule fires the moment the condition is met (this minimizes the waiting time for the demonstration.).
1. Set **Keep firing for** to, `0s`{{copy}}, so the alert stops firing immediately after the condition is no longer true.

# Configure notifications

- Select a **Contact point**. If you don’t have any contact points, add a [Contact point](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/#add-a-contact-point).

  For a quick test, you can use a public webhook from [webhook.site](https://webhook.site/) to capture and inspect alert notifications. If you choose this method, select **Webhook** from the drop-down menu in contact points.

# Configure notification message

To link this alert rule to our visualization click [**Link dashboard and panel**](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/link-alert-rules-to-panels/#link-alert-rules-to-panels)

- Select the folder that contains the dashboard. In this case: **system-metrics**
- Select the **cpu-and-memory-metrics** visualization
- Click **confirm**

You have successfully linked this alert rule to your visualization!

When the CPU usage exceeds the defined threshold, an annotation should appear on the graph to mark the event. Similarly, when the alert is resolved, another annotation is added to indicate the moment it returned to normal.
