1. In **Folder**, click _+ New folder_ and enter a name. For example: `app-metrics`{{copy}} . This folder contains our alerts.
1. Click **+ Add labels**.
1. **Key** field: `environment`{{copy}} .
1. In the **value** field copy in the following template:

   ```go
   {{- $env := reReplaceAll ".*([pP]rod|[sS]taging|[dD]ev).*" "${1}" $labels.instance -}}
   {{- if eq $env "prod" -}}
   production
   {{- else if eq $env "staging" -}}
   staging
   {{- else -}}
   development
   {{- end -}}
   ```{{copy}}

   This template uses a regular expression to extract `prod`{{copy}}, `staging`{{copy}}, or `dev`{{copy}} from the instance label (`$labels.instance`{{copy}}) and maps it to a more readable label (like "production" for "prod").

As result, when alerts exceed a threshold, the template checks the labels, such as `instance="flask-prod:5000"`{{copy}}, `instance="flask-staging:5000"`{{copy}}, or custom labels like `deployment="prod-us-cs30"`{{copy}}, and assigns a value of production, staging or development to the custom environment **environment** label.

This label is then used by the alert notification policy to route alerts to the appropriate team, so that notifications are delivered efficiently, and reducing unnecessary noise.

# Set evaluation behaviour

1. Click + **New evaluation group**. Name it `system-usage`{{copy}}.
1. Choose an **Evaluation interval** (how often the alert will be evaluated). Choose `1m`{{copy}}.
1. Set the **pending period** to `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met (this minimizes the waiting time for the demonstration.).
1. Set **Keep firing for** to, `0s`{{copy}}, so the alert stops firing immediately after the condition is no longer true.

# Configure notifications

Select who should receive a notification when an alert rule fires.

1. Toggle the **Advance options** button.
1. Click **Preview routing**.
   The preview should display which firing alerts are routed to contact points based on notification policies that match the `environment`{{copy}} label.

   ![Notification policies matched by the environment label matcher.](https://grafana.com/media/docs/alerting/dynamic-routing-preview-prod-staging.png)

   The environment label matcher should map to the notification policies created earlier. This makes sure that firing alert instances are routed to the appropriate contact points associated with each policy.

 Step 3: Create a second alert rule for memory usage

1. Duplicate the existing alert rule (**More > Duplicate**), or create a new alert rule for memory usage, defining a threshold condition (e.g., memory usage exceeding `60%`{{copy}}).
1. Give it a name. For example: `memory-usage`{{copy}}
1. Query: `flask_app_memory_usage{}`{{copy}}
1. Link to the same visualization to obtain memory usage annotations whenever the alert rule triggers or resolves.

Now that the CPU and memory alert rules are set up, they are linked to the notification policies through the custom label matcher we added. The value of the label dynamically changes based on the environment template, using `$labels.instance`{{copy}}. This ensures that the label value will be set to production, staging, or development, depending on the environment.
