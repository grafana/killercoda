# Setting up alert rule grouping

## Notification Policy

Following the above example, [notification policies](ref:notification-policies) are created to route alert instances, which have a region label, to a specific contact point. The goal is to receive one consolidated notification per region. To demonstrate how grouping works, alert notifications for the East Coast team are not grouped. Regarding timing, a specific schedule is defined for that region. This setup overrides the parent's settings to fine-tune the behavior for specific labels (i.e., regions).

1. Visit [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana should be running
1. Navigate to **Alerts & IRM > Alerting > Notification policies**.
1. In the Default policy, click **+ New child policy**.

   - In the Default policy, click **+ New child policy**.
   - **Label**: `region`{{copy}}
   - **Operator**: `=`{{copy}}
   - **Value**: `us-west`{{copy}}

   This label matches alert rules where the region label is us-west
1. Choose a **Contact point**:

   - Select **Webhook**.

   If you don’t have any contact points, add a Contact point.
1. Enable Continue matching:

   - Turn on **Continue matching subsequent sibling nodes** so the evaluation continues even after one or more labels (i.e. region label) match.
1. Override grouping settings:

   - Toggle **Override grouping**.
   - **Group by**: `region`{{copy}}.

     **Group by** consolidates alerts that share the same grouping label into a single notification. For example, all alerts with `region=us-west`{{copy}} will be combined into one notification, making it easier to manage and reducing alert fatigue.
1. Set custom timing:

   - Toggle **Override general timings**.
   - **Group interval**: `2m`{{copy}}. This ensures follow-up notifications for the same alert group will be sent at intervals of 2 minutes. While the default is 5 minutes, we chose 2 minutes here to provide faster feedback for demonstration purposes.

     **Timing options** control how often notifications are sent and can help balance timely alerting with minimizing noise.
1. Save and repeat:

   - Repeat for `region = us-east`{{copy}} with a different webhook or a different contact point.

     **Note**: In Grafana, each label within a notification policy must have a unique key. If you attempt to add the same label key (e.g., region) with different values (us-west and us-east), only the last entry is saved, and the previous one is discarded. This is because labels are stored as associative arrays (maps), where each key must be unique.
     For identical label keys use regex matchers (e.g., _region=~"us-west|us-east"_).
