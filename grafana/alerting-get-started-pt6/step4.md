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
   flask_app_cpu_usage{instance="flask-prod:5000"}
   ```{{copy}}
1. **Alert condition**

   - Enter 75 as the value for **WHEN QUERY IS ABOVE** to set the threshold for the alert.
   - Click **Preview alert rule condition** to run the queries.

     ![Preview of a query returning alert instances in Grafana.](https://grafana.com/media/docs/alerting/alert-condition-details-prod.png)

   The query returns the CPU usage of the Flask application in the production environment. In this case, the usage is `86.01%`{{copy}}, which exceeds the configured threshold of `75%`{{copy}}, causing the alert to fire.
