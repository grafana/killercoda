# Setting up alert rule grouping

## Notification Policy

[Notification policies](ref:notification-policies) group alert instances and route notifications to specific contact points.

To follow the above example, we will create notification policies that route alert instances based on the `region`{{copy}} label to specific contact points. This setup ensures that alerts for a given region are consolidated into a single notification. Additionally, we will fine-tune the **timing settings** for each region by overriding the default parent policy, allowing more granular control over when notifications are sent.
