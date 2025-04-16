# Create Notification Policies

Notification policies route alert instances to contact points via label matchers. Since we know what labels our application returns (i.e., `environment`{{copy}}, `job`{{copy}}, `instance`{{copy}}), we can use these labels to match alert rules.

1. Navigate to **Alerts & IRM > Alerting > Notification Policies**.
1. Add a child policy:

   - In the **Default policy**, click **+ New child policy**.
   - **Label**: `environment`{{copy}}.
   - **Operator**: `=`{{copy}}.
   - **Value**: `production`{{copy}}.
   - This label matches alert rules where the environment label is `prod`{{copy}}.
1. Choose a **contact point**:

   - If you don’t have any contact points, add a [Contact point](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/#add-a-contact-point).

   For a quick test, you can use a public webhook from [webhook.site](https://webhook.site/) to capture and inspect alert notifications. If you choose this method, select **Webhook** from the drop-down menu in contact points.
1. Enable continue matching:

   - Turn on **Continue matching subsequent sibling nodes** so the evaluation continues even after one or more labels (i.e., _environment_ labels) match.
1. Save and repeat

   - Create another child policy by following the same steps.
   - Use `environment = staging`{{copy}} as the label/value pair.
   - Feel free to use a different contact point.

Now that the labels are defined, we can create alert rules for CPU and memory metrics. These alert rules will use the labels from the collected and stored metrics in Prometheus.
