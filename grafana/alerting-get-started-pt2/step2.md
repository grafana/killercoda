# Alert instances

An [alert instance](https://grafana.com/docs/grafana/latest/alerting/fundamentals/#alert-instances) is an event that matches a metric returned by an alert rule query.

Let's consider a scenario where you're monitoring website traffic using Grafana. You've set up an alert rule to trigger an alert instance if the number of page views exceeds a certain threshold (more than `1000`{{copy}} page views) within a specific time period, say, over the past `5`{{copy}} minutes.

If the query returns more than one time-series, each time-series represents a different metric or aspect being monitored. In this case, the alert rule is applied individually to each time-series.

![Screenshot displaying alert instances in the context of an alert rule, highlighting the specific alerts triggered by the rule and their respective statuses](https://grafana.com/media/docs/alerting/alert-instance-flow.jpg)

In this scenario, each time-series is evaluated independently against the alert rule. It results in the creation of an alert instance for each time-series. The time-series corresponding to the desktop page views meets the threshold and, therefore, results in an alert instance in **Firing** state for which an alert notification is sent. The mobile alert instance state remains **Normal**.
