# Create a Grafana-managed alert rule

Alert rules allow you to identify problems in your system moments after they occur. By quickly identifying unintended changes in your system, you can minimize disruptions to your services.

Grafana’s new alerting platform debuted with Grafana 8. A year later, with Grafana 9, it became the default alerting method. In this step we will create a Grafana-managed alert rule. Then we will trigger our new alert rule and send a test message to a dummy endpoint.

The most basic alert rule consists of two parts:

1. A _Contact point_ - A Contact point defines how Grafana delivers an alert instance. When the conditions of an _alert rule_ are met, Grafana notifies the contact points, or channels, configured for that alert rule.

> An [alert instance](https://grafana.com/docs/grafana/latest/alerting/fundamentals/#alert-instances) is a specific occurrence that matches a condition defined by an alert rule, such as when the rate of requests for a specific route suddenly increases.

<!-- raw HTML omitted -->

> An [alert instance](https://grafana.com/docs/grafana/latest/alerting/fundamentals/#alert-instances) is a specific occurrence that matches a condition defined by an alert rule, such as when the rate of requests for a specific route suddenly increases.

Some popular channels include:

- [Email](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-email/)

- [Webhooks](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/)

- [Telegram](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-telegram/)

- [Slack](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-slack/)

- [PagerDuty](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/pager-duty/)

1. An _Alert rule_ - An Alert rule defines one or more _conditions_ that Grafana regularly evaluates. When these evaluations meet the rule’s criteria, the alert rule is triggered.

To begin, let’s set up a webhook contact point. Once we have a usable endpoint, we’ll write an alert rule and trigger a notification.
