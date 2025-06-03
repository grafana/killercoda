# Visualizing metrics and alert annotations

![Time series panel displaying health indicators and annotations.](https://grafana.com/media/docs/alerting/panel-2-queries-and-alerts.png)

After the alert rules are created, they should appear as **health indicators** (colored heart icons: a red heart when the alert is in **Alerting** state, and a green heart when in **Normal** state) on the linked panel. In addition, annotations provide helpful context, such as the time the alert was triggered.

Finally, as part of the alerting process, you should receive notifications at the associated contact point.

```
{
  "receiver": "prod-alerts",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "cpu-usage",
        "deployment": "prod-us-cs30",
        "grafana_folder": "sys-metrics",
        "instance": "flask-prod:5000",
        "job": "flask"
      },
      "annotations": {},
      "silenceURL": "http://localhost:3000/alerting/silence/new?
      "dashboardURL": "http://localhost:3000/d/dc203378-1ef9-410b-a636-b533a0dd3bd8?from=1748934450000&orgId=1&to=1748938080006",
      "panelURL": "http://localhost:3000/d/dc203378-1ef9-410b-a636-b533a0dd3bd8?from=1748934450000&orgId=1&to=1748938080006&viewPanel=2",

... }
```{{copy}}

_Received alert notification in webhook Contact point_

It’s worth mentioning that alert rules that are linked to a panel include a link to said visualization in the alert notifications. In the alert notification example above, the message includes useful information such as the summary, description, and a link to the relevant dashboard for the firing or resolved alert (i.e. `dashboardURL`{{copy}}). This helps responders quickly navigate to the appropriate context for investigation.

You can extend this functionality by adding a custom annotation to your alert rules and creating a notification template that [includes a link to a dashboard with a time range.](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/template-notifications/examples/#print-a-link-to-a-dashboard-with-time-range). The URL will include a time range based on the alert’s timing—starting from one hour before the alert started (`from`{{copy}}) to either the alert’s end time or the current time (`to`{{copy}}), depending on whether the alert is resolved or still firing.

The final URL is constructed using a custom annotation (e.g., `MyDashboardURL`{{copy}}) along with the `from`{{copy}} and `to`{{copy}} parameters, which are calculated in the notification template.
