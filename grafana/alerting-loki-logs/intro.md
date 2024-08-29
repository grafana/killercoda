# How to create alert rules with log data

Loki stores your logs and only indexes labels for each log stream. Using Loki with Grafana Alerting is a powerful way to keep track of what’s happening in your environment. You can create metric alert rules based on content in your log lines to notify your team. What’s even better is that you can add label data from the log message directly into your alert notification.

In this tutorial, you’ll:

- Generate sample logs and pull them with Promtail to Grafana.

- Create an alert rule based on a Loki query (LogQL).

- Create a Webhook contact point to send alert notifications to.

> Check out our advanced alerting tutorial to explore advanced topics such as alert instances and notification routing.
