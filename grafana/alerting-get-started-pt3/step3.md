# A real-world example of alert grouping in action

## Scenario: monitoring a distributed application

You’re monitoring metrics like CPU usage, memory utilization, and network latency across multiple regions. Some of these alert rules include labels such as `region: us-west`{{copy}} and `region: us-east`{{copy}}. If multiple alert rules trigger across these regions, they can result in notification floods.

## How to manage grouping

To group alert rule notifications:

1. **Define labels**: Use `region`{{copy}}, `metric`{{copy}}, or `instance`{{copy}} labels to categorize alerts.

1. **Configure Notification policies**:
   - Group alerts by the **query label** “region”.

   - Example:
     - Alert notifications for `region: us-west`{{copy}} go to the West Coast team.

     - Alert notifications for `region: us-east`{{copy}} go to the East Coast team.

1. Specify the **timing options** for sending notifications to control their frequency.
   - Example:
     - **Group interval**: setting determines how often updates for the same alert group are sent. By default, this interval is set to 5 minutes, but you can customize it to be shorter or longer based on your needs.
