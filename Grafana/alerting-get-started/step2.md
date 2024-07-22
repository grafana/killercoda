# Create a contact point

Besides being an open-source observability tool, Grafana has its own built-in alerting service. This means that you can receive notifications whenever there is an event of interest in your data, and even see these events graphed in your visualizations.

In this step, we’ll set up a new [contact point](https://grafana.com/docs/grafana/latest/alerting/configure-notifications/manage-contact-points/integrations/webhook-notifier/). This contact point will use the _webhooks_ integration. In order to make this work, we also need an endpoint for our webhook integration to receive the alert. We will use [Webhook.site](https://webhook.site/) to quickly set up that test endpoint. This way we can make sure that our alert is actually sending a notification somewhere.

1. In your browser, **sign in** to your Grafana Cloud account.

   {{< admonition type=“note” >}}
   **OSS users**: To log in, navigate to localhost:3000, where Grafana is running locally.
   {{< /admonition >}}

1. In another tab, go to [Webhook.site](https://webhook.site/).

1. **Copy Your unique URL**.

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

# Create an alert

Next, we’ll establish an [alert rule](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/) within Grafana Alerting to notify us whenever our sample app experiences a specific volume of requests.

In Grafana, **navigate to Alerting** > **Alert rules**. Click on **New alert rule**.

## Enter alert rule name

Make it short and descriptive as this will appear in your alert notification. For instance, **database-metrics**

## Define query and alert condition

In this section, we define queries, expressions (used to manipulate the data), and the condition that must be met for the alert to be triggered.

1. Select the **Prometheus** data source from the drop-down menu.

1. In the Query editor, switch to **Code** mode by clicking the button at the right.

1. Enter the following query:

   ```promql
   vector(1)
   ```{{copy}}

   In Prometheus, `vector(1)`{{copy}} is a special type of PromQL query that generates a constant vector. This is useful in testing and query manipulation, where you might need a constant value for calculations or comparisons. This query will allow you to create an alert rule that will be always firing.

1. Remove the ‘B’ **Reduce expression** (click the bin icon). The Reduce expression comes by default, and in this case, it is not needed since the queried data is already reduced. Note that the Threshold expression is now your **Alert condition**.

1. In the ‘C’ **Threshold expression**:

   - Change the **Input** to **‘A’** to select the data source.

   - Enter `0`{{copy}} as the threshold value. This is the value above which the alert rule should trigger.

1. Click **Preview** to run the queries.

   It should return a single sample with the value 1 at the current timestamp. And, since `1`{{copy}} is above `0`{{copy}}, the alert condition has been met, and the alert rule state is `Firing`{{copy}}.

   ![A preview of a firing alert](https://grafana.com/media/docs/alerting/alerting-always-firing-alert.png)

## Set evaluation behavior

An [evaluation group](https://grafana.com/docs/grafana/latest/alerting/fundamentals/alert-rules/rule-evaluation/) defines when an alert rule fires, and it’s based on two settings:

- **Evaluation group**: how frequently the alert rule is evaluated.

- **Evaluation interval**: how long the condition must be met to start firing. This allows your data time to stabilize before triggering an alert, helping to reduce the frequency of unnecessary notifications.

To set up the evaluation:

1. In **Folder**, click **+ New folder** and enter a name. For example: _metric-alerts_. This folder will contain our alerts.

1. In the **Evaluation group**, repeat the above step to create a new evaluation group. We will name it _1m-evaluation_.

1. Choose an **Evaluation interval** (how often the alert will be evaluated).
   For example, every `1m`{{copy}} (1 minute).

1. Set the pending period to, `0s`{{copy}} (zero seconds), so the alert rule fires the moment the condition is met.

## Configure labels and notifications

Choose the contact point where you want to receive your alert notifications.

1. Under **Contact point**, select **Webhook** from the drop-down menu.

1. Click **Save rule and exit** at the top right corner.

# Trigger and resolve an alert

Now that the alert rule has been configured, you should receive alert [notifications](http://grafana.com/docs/grafana/next/alerting/fundamentals/alert-rule-evaluation/state-and-health/#notifications) in the contact point whenever alerts trigger and get resolved.

## Trigger an alert

Since the alert rule that you have created has been configured to always fire, once the evaluation interval has concluded, you should receive an alert notification in the Webhook endpoint.

![Firing alert notification details](https://grafana.com/media/docs/alerting/alerting-webhook-firing-alert.png)

The alert notification details show that the alert rule state is Firing , and it includes the value that made the rule trigger by exceeding the threshold of the alert rule condition. The notification also includes links to see the alert rule details, and another link to add a [Silence](http://grafana.com/docs/grafana/next/alerting/configure-notifications/create-silence/) to it.

## Resolve an alert

To see how a resolved alert notification looks like, you can modify the current alert rule threshold.

To edit the Alert rule:

1. **Navigate to Alerting** > **Alert rules**.

1. Click on the metric-alerts folder to display the alert that you created earlier

1. Click the **edit** button on the right hand side of the screen

1. Increment the Threshold expression to 1.

1. Click **Save rule and exit**.

By incrementing the threshold, the condition is no longer met, and after the evaluation interval has concluded (1 minute approx.), you should receive an alert notification with status **“Resolved”**.

# Learn more

Your learning journey continues in [Part 2](http://grafana.com/tutorials/alerting-get-started-pt2/) where you will learn about alert instances and notification routing.
