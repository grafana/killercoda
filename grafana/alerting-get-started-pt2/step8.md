# Receive alert notifications

Now that the alert rule has been configured, you should receive alert [notifications](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/state-and-health/#notifications) in the contact point whenever the alert triggers and gets resolved. In our example, each alert instance should be routed separately as we configured labels to match notification policies. Once the evaluation interval has concluded (1m), you should receive an alert notification in the Webhook endpoint.

![Screenshot showing the exploration of alert notification details in a webhook endpoint, displaying the content and structure of the alert payload received by the endpoint](https://grafana.com/media/docs/alerting/get-started-webhook-alert-isntance.png)

The alert notification details show that the alert instance corresponding to the website views from desktop devices was correctly routed through the notification policy to the Webhook contact point. The notification also shows that the instance is in **Firing** state, as well as it includes the label `device=desktop`{{copy}}, which makes the routing of the alert instance possible.

Feel free to change the CSV data in the alert rule to trigger the routing of the alert instance that matches the label `device=mobile`{{copy}}.
