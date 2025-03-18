# Notification policies

[Notification policies](https://grafana.com/docs/grafana/latest/alerting/fundamentals/notifications/notification-policies/) route alerts to different communication channels, reducing alert noise and providing control over when and how alerts are sent. For example, you might use notification policies to ensure that critical alerts about server downtime are sent immediately to the on-call engineer. Another use case could be routing performance alerts to the development team for review and action.

Key Characteristics:

- Route alert notifications by matching alerts and policies with labels
- Manage when to send notifications

![Screenshot illustrating the routing of alerts with notification policies, including the configuration and flow of alerts through different notification channels](https://grafana.com/media/docs/alerting/get-started-notification-policy-tree-combo.png)

In the above diagram, alert instances and notification policies are matched by labels. For instance, the label `team=operations`{{copy}} matches the alert instance “**Pod stuck in CrashLoop**” and “**Disk Usage -80%**” to child policies that send alert notifications to a particular contact point (<operations@grafana.com>).
