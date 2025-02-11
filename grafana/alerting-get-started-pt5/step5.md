# Create alert rules to monitor CPU and memory usage

Follow these steps to manually create alert rules and link them to a visualization.

# Create an alert rule for CPU usage

1. Navigate to **Alerts & IRM > Alerting > Alert rules** from the Grafana sidebar.

1. Click **+ New alert** rule to create a new alert.

## Enter alert rule name

Make it short and descriptive as this will appear in your alert notification. For instance, `CPU usage`{{copy}} .

## Define query and alert condition

1. Select **Prometheus** data source from the drop-down menu.

1. In the query section, enter the following query:

   ** switch to **Code** mode if not already selected **

   ```
   flask_app_cpu_usage{}
   ```{{copy}}

1. **Alert condition** section:

   - Enter 75 as the value for **WHEN QUERY IS ABOVE** to set the threshold for the alert.

   - Click **Preview alert rule condition** to run the queries.

   ![Preview of a query returning alert instances in Grafana.](https://grafana.com/media/docs/alerting/promql-returning-metrics.png)

   Among the labels returned for `flask_app_cpu_usage`{{copy}}, the environment label is particularly important, as it enables dynamic alert routing based on the environment value, ensuring the right team receives the relevant notifications.

## Add folders and labels

In this section we add a [templated label based on query value](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/templates/examples/#based-on-query-value) to map to the notification policies.

1. In **Folder**, click _+ New folder_ and enter a name. For example: `App metrics`{{copy}} . This folder contains our alerts.

1. Click **+ Add labels**.

1. **Key** field: `environment`{{copy}} .

1. In the **value** field copy in the following template:

   ```go
   {{- if eq $labels.environment "prod" -}}
   production
   {{- else if eq $labels.environment "staging" -}}
   staging
   {{- else -}}
   development
   {{- end -}}
   ```{{copy}}

   In this context, the template is used to route alert notifications based on the `environment`{{copy}} label. When a metric like CPU usage exceeds a threshold, the template checks the environment (e.g., `prod`{{copy}}, `staging`{{copy}}, or any other value). It then generates a label based on query value (e.g., _production_, _staging_, or _development_). This label is used in the alert notification policy to route alerts to the appropriate team, so that notifications are directed to the right group, making the process more efficient and avoiding unnecessary overlap.

## Set evaluation behaviour

1. Click + **New evaluation group**. Name it `System usage`{{copy}}.

1. Choose an **Evaluation interval** (how often the alert will be evaluated). Choose `1m`{{copy}}. Click Create.

1. Set the **pending period** to `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met (this minimizes the waiting time for the demonstration.).

## Configure notifications

Select who should receive a notification when an alert rule fires.

1. Toggle the **Advance options** button.

1. Click **Preview routing**.
   The preview should display which firing alerts are routed to contact points based on notification policies that match the `environment`{{copy}} label.

   ![Notification policies matched by the environment label matcher.](https://grafana.com/media/docs/alerting/routing-preview-cpu-metrics.png)

   The environment label matcher should map to the notification policies created earlier. This makes sure that firing alert instances are routed to the appropriate contact points associated with each policy.

## Configure notification message

Link your dashboard panel to this alert rule to display alert annotations in your visualization whenever the alert rule triggers or resolves.

1. Click **Link dashboard and panel**.

1. Find the panel that you created earlier.

1. Click **Confirm**.

# Create a second alert rule for memory usage

1. Duplicate the existing alert rule (**More > Duplicate**), or create a new alert rule for memory usage, defining a threshold condition (e.g., memory usage exceeding `60%`{{copy}}).

1. Query: `flask_app_memory_usage{}`{{copy}}

1. Link to the same visualization to obtain memory usage annotations whenever the alert rule triggers or resolves.

Now that the CPU and memory alert rules are set up, they are linked to the notification policies through the custom label matcher we added. The value of the label dynamically changes based on the environment template, using `$labels.environment`{{copy}}. This ensures that the label value will be set to production, staging, or development, depending on the environment.
