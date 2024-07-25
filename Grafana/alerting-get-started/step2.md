# Create a contact point

Besides being an open-source observability tool, Grafana has its own built-in alerting service. This means that you can receive notifications whenever there is an event of interest in your data, and even see these events graphed in your visualizations.

In this step, we’ll set up a new [contact point](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/). This contact point will use the _webhooks_ integration. In order to make this work, we also need an endpoint for our webhook integration to receive the alert. We will use [Webhook.site](https://webhook.site/) to quickly set up that test endpoint. This way we can make sure that our alert is actually sending a notification somewhere.

1. In your browser, **sign in** to your Grafana Cloud account.

   OSS users: To log in, navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana is running.

1. In another tab, go to [Webhook.site](https://webhook.site/).

1. Copy Your unique URL.

Your webhook endpoint is now waiting for the first request.

Next, let’s configure a contact point in Grafana’s Alerting UI to send notifications to our webhook endpoint.

1. Return to Grafana. In Grafana’s sidebar, hover over the **Alerting** (bell) icon and then click **Contact points**.

1. Click **+ Add contact point**.

1. In **Name**, write **Webhook**.

1. In **Integration**, choose **Webhook**.

1. In **URL**, paste the endpoint to your webhook endpoint.

1. Click **Test**, and then click **Send test notification** to send a test alert to your webhook endpoint.

1. Navigate back to [Webhook.site](https://webhook.site/). On the left side, there’s now a `POST /`{{copy}} entry. Click it to see what information Grafana sent.

   ![A POST entry in Webhook.site](https://grafana.com/media/docs/alerting/alerting-webhook-detail.png)

1. Return to Grafana and click **Save contact point**.

We have created a dummy Webhook endpoint and created a new Alerting contact point in Grafana. Now, we can create an alert rule and link it to this new integration.
