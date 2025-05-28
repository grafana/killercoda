# Step 2: Create alert rules to monitor CPU and memory usage

Follow these steps to manually create alert rules and link them to a visualization.

# Create an alert rule for CPU usage

1. Navigate to **Alerts & IRM > Alerting > Alert rules** from the Grafana sidebar.
1. Click **+ New alert rule** rule to create a new alert.

## Enter alert rule name

Make it short and descriptive, as this will appear in your alert notification. For instance, `cpu-usage`{{copy}} .

## Define query and alert condition

1. Select **Prometheus** data source from the drop-down menu.
1. In the query section, enter the following query:

   ** switch to **Code** mode if not already selected **

   ```
   flask_app_cpu_usage{}
   ```{{copy}}
1. **Alert condition** section:

   - Enter `75`{{copy}} as the value for **WHEN QUERY IS ABOVE** to set the threshold for the alert.
   - Click **Preview alert rule condition** to run the queries.

   ![Preview of a query returning alert instances in Grafana.](https://grafana.com/media/docs/alerting/flask-app-metrics.png)

   Among the labels returned for `flask_app_cpu_usage`{{copy}}, the labels `instance`{{copy}} and `deployment`{{copy}} contain values that include the term _prod_ and _staging_. We will create a template later to detect these keywords, so that any firing alert instances are routed to the relevant contact points (e.g., alerts-prod, alerts-staging).

## Add folders and labels

In this section we add a [templated label based on query value](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/templates/examples/#based-on-query-value) to map to the notification policies.
