# Create a Grafana Managed Alert

Alerts allow you to identify problems in your system moments after they occur. By quickly identifying unintended changes in your system, you can minimize disruptions to your services.

Grafana’s new alerting platform debuted with Grafana 8. A year later, with Grafana 9, it became the default alerting method. In this step we will create a Grafana Managed Alert. Then we will trigger our new alert and send a test message to a dummy endpoint.

The most basic alert consists of two parts:

1. A _Contact point_ - A Contact point defines how Grafana delivers an alert. When the conditions of an _alert rule_ are met, Grafana notifies the contact points, or channels, configured for that alert.

   Some popular channels include:

   - [Email](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-email/)

   - [Webhooks](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/)

   - [Telegram](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-telegram/)

   - [Slack](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/configure-slack/)

   - [PagerDuty](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/pager-duty/)

1. An _Alert rule_ - An Alert rule defines one or more _conditions_ that Grafana regularly evaluates. When these evaluations meet the rule’s criteria, the alert is triggered.

To begin, let’s set up a webhook contact point. Once we have a usable endpoint, we’ll write an alert rule and trigger a notification.
