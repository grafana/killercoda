# Step 4: Receiving notifications

Finally, as part of the alerting process, you should receive notifications at the associated contact point. If you're receiving alerts via email, the default email template will include two buttons:

- **View dashboard**: links to the dashboard that contains the alerting panel
- **View panel**: links directly to the individual panel where the alert was triggered

![Alert notification with links to panel and dashboard.](https://grafana.com/media/docs/alerting/email-notification-w-url.png)

Clicking either button opens Grafana with a pre-applied time range relevant to the alert.

By default, this URL includes `from`{{copy}} and `to`{{copy}} query [parameters](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/template-notifications/reference/#alert) that reflect the time window around the alert event (one hour before and after the alert). This helps you land directly in the time window where the alert occurred, making it easier to analyze what happened.

If you want to define a more intentional time range, you can customize your notifications using a [notification template](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/template-notifications/examples/#print-a-link-to-a-dashboard-with-time-range). With a template, you can explicitly set `from`{{copy}} and `to`{{copy}} values for more precise control over what users see when they follow the dashboard link. The final URL is constructed using a custom annotation (e.g., `MyDashboardURL`{{copy}}) along with the `from`{{copy}} and `to`{{copy}} parameters, which are calculated in the notification template.
