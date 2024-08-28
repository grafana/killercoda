# Create a contact point for Grafana-managed alert rules

In this step, we’ll set up a new contact point. This contact point will use the _webhooks_ channel. In order to make this work, we also need an endpoint for our webhook channel to receive the alert notification. We will use [requestbin.com](https://requestbin.com) to quickly set up that test endpoint. This way we can make sure that our alert manager is actually sending a notification somewhere.

1. Browse to [requestbin.com](https://requestbin.com).

1. Under the **Create Request Bin** button, click the link to create a **public bin** instead.

1. From Request Bin, copy the endpoint URL.

Your Request Bin is now waiting for the first request.

Next, let’s configure a Contact Point in Grafana’s Alerting UI to send notifications to our Request Bin.

1. Return to Grafana. In Grafana’s sidebar, hover over the **Alerting** (bell) icon and then click **Contact points**.

1. Click **+ Add contact point**.

1. In **Name**, write **RequestBin**.

1. In **Integration**, choose **Webhook**.

1. In **URL**, paste the endpoint to your request bin.

1. Click **Test**, and then click **Send test notification** to send a test alert notification to your request bin.

1. Navigate back to the Request Bin you created earlier. On the left side, there’s now a `POST /`{{copy}} entry. Click it to see what information Grafana sent.

1. Return to Grafana and click **Save contact point**.

We have now created a dummy webhook endpoint and created a new Alerting Contact Point in Grafana. Now we can create an alert rule and link it to this new channel.
