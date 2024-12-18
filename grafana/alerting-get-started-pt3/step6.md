# Receiving grouped alert notifications

Now that the alert rule has been configured, you should receive alert notifications in the contact point whenever alerts trigger.

When the configured alert rule detects CPU usage higher than 75% across multiple regions, it will evaluate the metric every minute. If the condition persists, notifications will be grouped together, with a **Group wait** of 30 seconds before the first alert is sent. Follow-up notifications are sent every 2 minutes for quick updates in this demonstration, but for reducing alert frequency, consider using the default or increasing the interval. If the condition continues for an extended period, a **Repeat interval** of 4 hours ensures that the alert is only resent if the issue persists

As a result, our notification policy will route two notifications: one notification grouping the three alert instances from the `us-east`{{copy}} region and another grouping the two alert instances from the `us-west`{{copy}} region

Grouped notifications example:

Webhook - US East

```json
{
  "receiver": "webhook-us-east",
  "status": "firing",
  "alerts": [{ "instance": "server-03" }, { "instance": "server-06" }, { "instance": "server-08" }]
}
```{{copy}}

Webhook - US West

```json
{
  "receiver": "webhook-us-west",
  "status": "firing",
  "alerts": [{ "instance": "server-02" }, { "instance": "server-07" }]
}
```{{copy}}
