# How to create alert rules with log data

Loki stores your logs and only indexes labels for each log stream. Using Loki with Grafana Alerting is a powerful way to keep track of what’s happening in your environment. You can create metric alert rules based on content in your log lines to notify your team. What’s even better is that you can add label data from the log message directly into your alert notification.

In this tutorial, you'll:

- Generate sample logs and collect them with Grafana Alloy.
- Create an alert rule based on a Loki query (LogQL).
- Create a Webhook contact point to send alert notifications to.

> In [Get started with Grafana Alerting - Part 2](https://grafana.com/tutorials/alerting-get-started-pt2/) you can advance your skills by exploring alert instances and notification routing.
