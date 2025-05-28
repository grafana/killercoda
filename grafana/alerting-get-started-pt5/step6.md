# Done! Your alerts are now dynamically routed

Based on your query's `instance`{{copy}} label values (which contain keywords like _prod_ or _staging_ ), Grafana dynamically assigns the value `production`{{copy}}, `staging`{{copy}} or `development`{{copy}} to the custom **environment** label using the template. This dynamic label then matches the label matchers in your notification policies, which route alerts to the correct contact points.

To see this in action go to **Alerts & IRM > Alerting > Active notifications**

This page shows grouped alerts that are currently triggering notifications. If you click on any alert group to view its label set, contact point, and number of alert instances. Notice that the **environment** label has been dynamically populated with values like `production`{{copy}}.

![Expanded alert in Active notifications section](https://grafana.com/media/docs/alerting/routing-active-notification-detail.png)

Finally, you should receive notifications at the contact point associated with either `prod`{{copy}} or `staging`{{copy}}.

Feel free to experiment by changing the template to match other labels that contain any of the watched keywords. For example, you could reference:

```go
$labels.deployment
```{{copy}}

The template should be flexible enough to capture the target keywords (e.g., prod, staging) by adjusting which label the[`$labels`{{copy}}](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/templates/reference/#labels) is referencing.
