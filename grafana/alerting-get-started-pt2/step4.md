# Create notification policies

Create a notification policy if you want to handle metrics returned by alert rules separately by routing each alert instance to a specific contact point. In Grafana, click on the icon at the top left corner of the screen to access the navigation menu.

1. Visit [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana should be running

1. Navigate to **Alerts & IRM > Alerting > Notification policies**.

1. In the Default policy, click **+ New child policy**.

1. In the field **Label** enter `device`{{copy}}, and in the field **Value** enter `desktop`{{copy}}.

1. From the **Contact point** drop-down, choose **Webhook**.

   If you don’t have any contact points, add a [Contact point](https://grafana.com/tutorials/alerting-get-started/#create-a-contact-point).

1. Click **Save Policy**.

   This new child policy routes alerts that match the label `device=desktop`{{copy}} to the Webhook contact point.

1. **Repeat the steps above to create a second child policy** to match another alert instance. For labels use: `device=mobile`{{copy}}. Use the Webhook integration for the contact point. Alternatively, experiment by using a different Webhook endpoint or a [different integration](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/#list-of-supported-integrations).
