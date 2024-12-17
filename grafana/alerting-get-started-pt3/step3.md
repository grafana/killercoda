# A real-world example of alert grouping in action

## Scenario: monitoring a distributed application

You’re monitoring metrics like CPU usage, memory utilization, and network latency across multiple regions. Alert rules include labels such as `region: us-west`{{copy}} and `region: us-east`{{copy}}. If multiple alerts trigger across these regions, they can result in notification floods.

## How to manage grouping

To group alert rule notifications:

1. **Define labels**: Use `region`{{copy}}, `metric`{{copy}}, or `instance`{{copy}} labels to categorize alerts.

1. **Configure Notification policies**:
   - Group alerts by the `region`{{copy}} label.

   - Example:
     - Alerts for `region: us-west`{{copy}} go to the West Coast team.

     - Alerts for `region: us-east`{{copy}} go to the East Coast team.
