# How templating works

In Grafana, you can use [templates](ref:templates) to dynamically pull in specific data about the alert rule. This results in more flexible and informative alert notification messages. You can template either alert rule labels and annotations, or notification templates. Both use the Go template language.

![How templating works](https://grafana.com/media/docs/alerting/how-notification-templates-works.png)

## Templating alert rule labels and annotations

[Labels and annotations](ref:template-labels-annotations) are key fields where templates are applied. One of the main advantages of using templating in annotations is the ability to incorporate dynamic data from queries, allowing alerts to reflect real-time information relevant to the triggered condition. By using templating in annotations, you can customize the content of each alert instance, such as including instance names and metric values, so the notification becomes more informative.

## Notification templates

The real power of templating lies in how it helps you format notifications with dynamic alert data. [Notification templates](ref:template-notifications) let you pull in details from annotations to create clear and consistent messages. They also make it simple to reuse the same format across different contact points, saving time and effort.

Notification templates allow you to customize how information is presented in each notification. For example, you can use templates to organize and format details about firing or resolved alerts, making it easier for recipients to understand the status of each alert at a glance—all within a single notification.

This particular notification template pulls in summary and description annotations for each alert instance and organizes them into separate sections, such as “firing” and “resolved.” This way, instead of getting a long list of individual alert notifications, users can receive one well-structured message with all the relevant details grouped together.

This approach is helpful when you want to reduce notification noise, especially in situations where multiple instances of an alert are firing at the same time (e.g., high CPU usage across several instances). You can leverage templates to create a unified, easy-to-read notification that includes all the pertinent details.
