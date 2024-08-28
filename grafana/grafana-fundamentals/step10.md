# Add an alert rule to Grafana

Now that Grafana knows how to notify us, it’s time to set up an alert rule:

1. In Grafana’s sidebar, hover over the **Alerting** (bell) icon and then click **Alert rules**.

1. Click **+ New alert rule**.

1. For **Section 1**, name the rule `fundamentals-test`{{copy}}.

1. For **Section 2**, Find the **query A** box. Choose your Prometheus datasource. Note that the rule type should automatically switch to Grafana-managed alert rule.

1. Switch to code mode by checking the Builder/Code toggle.

1. Enter the same Prometheus query that we used in our earlier panel:

   ```
   sum(rate(tns_request_duration_seconds_count[5m])) by(route)
   ```{{copy}}

1. Press **Preview**. You should see some data returned.

1. Keep expressions “B” and “C” as they are. These expressions (Reduce and Threshold, respectively) come by default when creating a new rule. Expression “B”, selects the last value of our query “A”, while the Threshold expression “C” will check if the last value from expression “B” is above a specific value. In addition, the Threshold expression is the alert rule condition by default. Enter `0.2`{{copy}} as threshold value. [You can read more about queries and conditions here](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/queries-conditions/#expression-queries).

1. In **Section 3**, in Folder, create a new folder, by clicking `New folder`{{copy}} and typing a name for the folder. This folder will contain our alert rules. For example: `fundamentals`{{copy}}. Then, click `create`{{copy}}.

1. In the Evaluation group, repeat the above step to create a new one. We will name it `fundamentals`{{copy}} too.

1. Choose an Evaluation interval (how often the alert rule will be evaluated). For example, every `10s`{{copy}} (10 seconds).

1. Set the pending period. This is the time that a condition has to be met until the alert instance enters in Firing state and a notification is sent. Enter `0s`{{copy}}. For the purposes of this tutorial, the evaluation interval is intentionally short. This makes it easier to test. This setting makes Grafana wait until an alert instance has fired for a given time before Grafana sends the notification.

1. In **Section 4**, choose **RequestBin** as the **Contact point**.

1. Click **Save rule and exit** at the top of the page.

# Trigger a Grafana-managed alert rule

We have now configured an alert rule and a contact point. Now let’s see if we can trigger a Grafana-managed alert rule by generating some traffic on our sample application.

1. Browse to [localhost:8081]({{TRAFFIC_HOST1_8081}}).

1. Add a new title and URL, repeatedly click the vote button, or refresh the page to generate a traffic spike.

Once the query `sum(rate(tns_request_duration_seconds_count[5m])) by(route)`{{copy}} returns a value greater than `0.2`{{copy}} Grafana will trigger our alert rule. Browse to the Request Bin we created earlier and find the sent Grafana alert notification with details and metadata.

Let’s see how we can configure this.

1. In Grafana’s sidebar, hover over the **Alerting** (bell) icon and then click **Alert rules**.

1. Expand the `fundamentals > fundamentals`{{copy}} folder to view our created alert rule.

1. Click the **Edit** icon and scroll down to **Section 5**.

1. Click the **Link dashboard and panel** button and select the dashboard and panel to which you want the alert instance to be added as an annotation.

1. Click **Confirm** and **Save rule and exit** to save all the changes.

1. In Grafana’s sidebar, navigate to the dashboard by clicking **Dashboards** and selecting the dashboard you created.

1. To test the changes, follow the steps listed to [trigger a Grafana-managed alert rule](https://grafana.com#trigger-a-grafana-managed-alert).

   You should now see a red, broken heart icon beside the panel name, signifying that the alert rule has been triggered. An annotation for the alert instance, represented as a vertical red line, is also displayed.

   ![A panel in a Grafana dashboard with alerting and annotations configured](https://grafana.com/media/tutorials/grafana-alert-on-dashboard.png)

> Check out our [advanced alerting tutorial](http://grafana.com/tutorials/alerting-get-started-pt2/) for more insights and tips.
