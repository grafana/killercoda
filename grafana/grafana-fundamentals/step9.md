# Create a contact point for Grafana-managed alert rules

In this step, we set up a new contact point. This contact point uses the _webhooks_ channel. In order to make this work, we also need an endpoint for our webhook channel to receive the alert notification. We can use [Webhook.site](https://webhook.site/) to quickly set up that test endpoint. This way we can make sure that our alert manager is actually sending a notification somewhere.

1. Browse to [Webhook.site](https://webhook.site/).
1. Copy Your unique URL.

Your webhook endpoint is now waiting for the first request.

Next, let's configure a Contact Point in Grafana's Alerting UI to send notifications to our webhook endpoint.

1. Return to Grafana. In Grafana's sidebar, hover over the **Alerting** (bell) icon and then click **Manage Contact points**.
1. Click **+ Add contact point**.
1. In **Name**, write **Webhook**.
1. In **Integration**, choose **Webhook**.
1. In **URL**, paste the endpoint to your webhook endpoint.
1. Click **Test**, and then click **Send test notification** to send a test alert notification to your webhook endpoint.
1. Navigate back to the webhook endpoint you created earlier. On the left side, there's now a `POST /`{{copy}} entry. Click it to see what information Grafana sent.
1. Return to Grafana and click **Save contact point**.

We have now created a dummy webhook endpoint and created a new Alerting Contact Point in Grafana. Now we can create an alert rule and link it to this new channel.
