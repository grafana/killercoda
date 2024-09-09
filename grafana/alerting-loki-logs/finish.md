<!-- raw HTML omitted -->

# How to create alert rules with log data

Loki stores your logs and only indexes labels for each log stream. Using Loki with Grafana Alerting is a powerful way to keep track of what’s happening in your environment. You can create metric alert rules based on content in your log lines to notify your team. What’s even better is that you can add label data from the log message directly into your alert notification.

In this tutorial, you’ll:

- Generate sample logs and pull them with Promtail to Grafana.

- Create an alert rule based on a Loki query (LogQL).

- Create a Webhook contact point to send alert notifications to.

> Check out our [advanced alerting tutorial](https://grafana.com/tutorials/alerting-get-started-pt2/) to explore advanced topics such as alert instances and notification routing.

<!-- raw HTML omitted -->

<!-- raw HTML omitted -->

## Before you begin

To demonstrate the observation of data using the Grafana stack, download the files to your local machine.

1. Download and save a Docker compose file to run Grafana, Loki and Promtail.

   ```bash
   wget https://raw.githubusercontent.com/grafana/loki/v2.8.0/production/docker-compose.yaml -O docker-compose.yaml
   ```{{exec}}

1. Run the Grafana stack.

   ```bash
   docker-compose up -d
   ```{{exec}}

The first time you run `docker-compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

> If you already have Grafana, Loki, or Prometheus running on your system, you might see errors, because the Docker image is trying to use ports that your local installations are already using. If this is the case, stop the services, then run the command again.

<!-- raw HTML omitted -->

<!-- raw HTML omitted -->

## Generate sample logs

1. Download and save a Python file that generates logs.

   ```bash
   wget https://raw.githubusercontent.com/grafana/tutorial-environment/master/app/loki/web-server-logs-simulator.py
   ```{{exec}}

1. Execute the log-generating Python script.

   ```bash
   python3 ./web-server-logs-simulator.py | sudo tee -a /var/log/web_requests.log
   ```{{exec}}

### Troubleshooting the script

If you don’t see the sample logs in Explore:

- Does the output file exist, check `/var/log/web_requests.log`{{copy}} to see if it contains logs.

- If the file is empty, check that you followed the steps above to create the file.

- If the file exists, verify that promtail container is running.

- In Grafana Explore, check that the time range is only for the last 5 minutes.

<!-- raw HTML omitted -->

<!-- raw HTML omitted -->

## Create a contact point

Besides being an open-source observability tool, Grafana has its own built-in alerting service. This means that you can receive notifications whenever there is an event of interest in your data, and even see these events graphed in your visualizations.

In this step, we’ll set up a new [contact point](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/). This contact point will use the _webhooks_ integration. In order to make this work, we also need an endpoint for our webhook integration to receive the alert. We will use [Webhook.site](https://webhook.site/) to quickly set up that test endpoint. This way we can make sure that our alert is actually sending a notification somewhere.

1. Navigate to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}), where Grafana is running.

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

<!-- raw HTML omitted -->

<!-- raw HTML omitted -->

## Create an alert rule

Next, we’ll establish an [alert rule](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/) within Grafana Alerting to notify us whenever alert rules are triggered and resolved.

1. In Grafana, **navigate to Alerting** > **Alert rules**.

1. Click on **New alert rule**.

1. Enter alert rule name for your alert rule. Make it short and descriptive as this will appear in your alert notification. For instance, **web-requests-logs**

### Define query and alert condition

In this section, we define queries, expressions (used to manipulate the data), and the condition that must be met for the alert to be triggered.

1. Select the **Loki** datasource from the drop-down.

1. In the Query editor, switch to Code mode by clicking the button on the right.

1. Paste the query below.

   ```
   sum by (message)(count_over_time({filename="/var/log/web_requests.log"} != "status=200" | pattern "<_> <message> duration<_>" [10m]))
   ```{{copy}}

This query will count the number of log lines with a status code that is not 200 (OK), then sum the result set by message type using an **instant query** and the time interval indicated in brackets. It uses the LogQL pattern parser to add a new label called `message`{{copy}} that contains the level, method, url, and status from the log line.

You can use the **explain query** toggle button for a full explanation of the query syntax. The optional log-generating script creates a sample log line similar to the one below:

```
2023-04-22T02:49:32.562825+00:00 level=info method=GET url=test.com status=200 duration=171ms
```{{copy}}

> Check out our [advanced alerting tutorial](https://grafana.com/tutorials/alerting-get-started-pt2/) to explore advanced topics such as alert instances and notification routing.

<!-- raw HTML omitted -->
